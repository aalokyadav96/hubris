package crops

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/crops/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewCropUseCase()).registerRoutes(router, app)
}

// Handler exposes crop HTTP endpoints.
type Handler struct {
	useCase usecase.CropUseCase
}

func NewHandler(uc usecase.CropUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmidware := middleware.Authenticate(app)
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/crops/health", h.health)
	router.HandlerFunc(http.MethodGet, "/api/v1/crops/crops", rateLimiter.Limit(GetCrops(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/crops/crops/count", rateLimiter.Limit(GetCropsCount(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/crops/crop", authmidware(CreateCrop(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/crops/crop/:cropid", GetCrop(app))
	router.HandlerFunc(http.MethodPut, "/api/v1/crops/crop/:cropid", authmidware(EditCrop(app)))
	router.HandlerFunc(http.MethodDelete, "/api/v1/crops/crop/:cropid", authmidware(DeleteCrop(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
