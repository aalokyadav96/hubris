package places

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/places/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewPlaceUseCase()).registerRoutes(router, app)
}

// Handler exposes place HTTP endpoints.
type Handler struct {
	useCase usecase.PlaceUseCase
}

func NewHandler(uc usecase.PlaceUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmidware := middleware.Authenticate(app)
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/places/health", h.health)
	router.HandlerFunc(http.MethodGet, "/api/v1/places/places", rateLimiter.Limit(GetPlaces(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/places/places/count", rateLimiter.Limit(GetPlacesCount(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/places/place", authmidware(CreatePlace(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/places/place/:placeid", GetPlace(app))
	router.HandlerFunc(http.MethodPut, "/api/v1/places/place/:placeid", authmidware(EditPlace(app)))
	router.HandlerFunc(http.MethodDelete, "/api/v1/places/place/:placeid", authmidware(DeletePlace(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
