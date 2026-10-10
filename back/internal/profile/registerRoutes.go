package profile

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/profile/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewProfileUseCase()).registerRoutes(router, app)
}

// Handler exposes profile HTTP endpoints.
type Handler struct {
	useCase usecase.ProfileUseCase
}

func NewHandler(uc usecase.ProfileUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmidware := middleware.Authenticate(app)
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/profiles/health", h.health)
	router.HandlerFunc(http.MethodGet, "/api/v1/profile/profiles", rateLimiter.Limit(GetProfiles(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/profile/profiles/count", rateLimiter.Limit(GetProfilesCount(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/profile/profile", authmidware(CreateProfile(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/profile/profile/:profileid", GetProfile(app))
	router.HandlerFunc(http.MethodPut, "/api/v1/profile/profile/:profileid", authmidware(EditProfile(app)))
	router.HandlerFunc(http.MethodDelete, "/api/v1/profile/profile/:profileid", authmidware(DeleteProfile(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
