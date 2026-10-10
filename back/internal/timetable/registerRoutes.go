package timetable

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/timetable/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewTimetableUseCase()).registerRoutes(router, app)
}

// Handler exposes timetable HTTP endpoints.
type Handler struct {
	useCase usecase.TimetableUseCase
}

func NewHandler(uc usecase.TimetableUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmidware := middleware.Authenticate(app)
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/timetable/health", h.health)
	router.HandlerFunc(http.MethodGet, "/api/v1/timetables/timetables", rateLimiter.Limit(GetTimetables(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/timetables/timetables/count", rateLimiter.Limit(GetTimetablesCount(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/timetables/timetable", authmidware(CreateTimetable(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/timetables/timetable/:timetableid", GetTimetable(app))
	router.HandlerFunc(http.MethodPut, "/api/v1/timetables/timetable/:timetableid", authmidware(EditTimetable(app)))
	router.HandlerFunc(http.MethodDelete, "/api/v1/timetables/timetable/:timetableid", authmidware(DeleteTimetable(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
