package events

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/events/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewEventUseCase()).registerRoutes(router, app)
}

// Handler exposes event HTTP endpoints.
type Handler struct {
	useCase usecase.EventUseCase
}

func NewHandler(uc usecase.EventUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmidware := middleware.Authenticate(app)
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/events/health", h.health)
	router.HandlerFunc(http.MethodGet, "/api/v1/events/events", rateLimiter.Limit(GetEvents(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/events/events/count", rateLimiter.Limit(GetEventsCount(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/events/event", authmidware(CreateEvent(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/events/event/:eventid", GetEvent(app))
	router.HandlerFunc(http.MethodPut, "/api/v1/events/event/:eventid", authmidware(EditEvent(app)))
	router.HandlerFunc(http.MethodDelete, "/api/v1/events/event/:eventid", authmidware(DeleteEvent(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
