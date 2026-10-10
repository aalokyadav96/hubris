package farms

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/farms/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewFarmUseCase()).registerRoutes(router, app)
}

// Handler exposes farm HTTP endpoints.
type Handler struct {
	useCase usecase.FarmUseCase
}

func NewHandler(uc usecase.FarmUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmidware := middleware.Authenticate(app)
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/farms/health", h.health)
	router.HandlerFunc(http.MethodGet, "/api/v1/farms/farms", rateLimiter.Limit(GetFarms(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/farms/farms/count", rateLimiter.Limit(GetFarmsCount(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/farms/farm", authmidware(CreateFarm(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/farms/farm/:farmid", GetFarm(app))
	router.HandlerFunc(http.MethodPut, "/api/v1/farms/farm/:farmid", authmidware(EditFarm(app)))
	router.HandlerFunc(http.MethodDelete, "/api/v1/farms/farm/:farmid", authmidware(DeleteFarm(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
