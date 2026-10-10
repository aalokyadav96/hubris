package tickets

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/tickets/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewTicketUseCase()).registerRoutes(router, app)
}

// Handler exposes ticket HTTP endpoints.
type Handler struct {
	useCase usecase.TicketUseCase
}

func NewHandler(uc usecase.TicketUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmidware := middleware.Authenticate(app)
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/tickets/health", h.health)
	router.HandlerFunc(http.MethodGet, "/api/v1/tickets/tickets", rateLimiter.Limit(GetTickets(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/tickets/tickets/count", rateLimiter.Limit(GetTicketsCount(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/tickets/ticket", authmidware(CreateTicket(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/tickets/ticket/:ticketid", GetTicket(app))
	router.HandlerFunc(http.MethodPut, "/api/v1/tickets/ticket/:ticketid", authmidware(EditTicket(app)))
	router.HandlerFunc(http.MethodDelete, "/api/v1/tickets/ticket/:ticketid", authmidware(DeleteTicket(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
