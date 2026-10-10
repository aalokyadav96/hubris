package timetable

import (
	"encoding/json"
	"net/http"

	"naevis/infra"
	"naevis/internal/timetable/usecase"

	"github.com/julienschmidt/httprouter"
)

// Handler exposes timetable HTTP endpoints.
type Handler struct {
	useCase usecase.TimetableUseCase
}

func NewHandler(uc usecase.TimetableUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	_ = app
	router.GET("/timetable/health", h.health)
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
