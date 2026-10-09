package places

import (
	"encoding/json"
	"net/http"
	"strconv"

	"naevis/infra"
	"naevis/infra/db"
	"naevis/internal/domain"
	"naevis/internal/usecase"
	"naevis/middleware"

	"github.com/julienschmidt/httprouter"
)

// Handler exposes place CRUD as HTTP handlers.
type Handler struct {
	useCase usecase.PlaceUseCase
}

func NewHandler(uc usecase.PlaceUseCase) *Handler {
	return &Handler{useCase: uc}
}

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	repo := db.NewPlaceRepository(app.PGPool)
	if repo == nil {
		router.GET("/places", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			http.Error(w, "database not configured", http.StatusServiceUnavailable)
		})
		return
	}

	h := NewHandler(usecase.NewPlaceUseCase(repo))

	router.GET("/places", wrap(app, h.listPlaces))
	router.POST("/places", wrap(app, h.createPlace))
	router.GET("/places/:id", wrap(app, h.getPlace))
	router.PUT("/places/:id", wrap(app, h.updatePlace))
	router.DELETE("/places/:id", wrap(app, h.deletePlace))
}

func wrap(app *infra.Deps, fn func(http.ResponseWriter, *http.Request, httprouter.Params)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authFn := middleware.Authenticate(app)(func(w http.ResponseWriter, r *http.Request) {
			fn(w, r, ps)
		})
		middleware.LoggingMiddleware(http.HandlerFunc(authFn)).ServeHTTP(w, r)
	}
}

func (h *Handler) listPlaces(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	places, err := h.useCase.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(places)
}

func (h *Handler) getPlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid place id", http.StatusBadRequest)
		return
	}
	place, err := h.useCase.GetByID(r.Context(), id)
	if err != nil {
		if err == domain.ErrPlaceNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(place)
}

func (h *Handler) createPlace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var place domain.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	created, err := h.useCase.Create(r.Context(), place)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}

func (h *Handler) updatePlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid place id", http.StatusBadRequest)
		return
	}
	var place domain.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	updated, err := h.useCase.Update(r.Context(), id, place)
	if err != nil {
		if err == domain.ErrPlaceNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(updated)
}

func (h *Handler) deletePlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid place id", http.StatusBadRequest)
		return
	}
	if err := h.useCase.Delete(r.Context(), id); err != nil {
		if err == domain.ErrPlaceNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
