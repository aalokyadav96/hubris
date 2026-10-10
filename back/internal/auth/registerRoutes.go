package auth

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/auth/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewAuthUseCase()).registerRoutes(router, app)
}

// Handler exposes auth HTTP endpoints.
type Handler struct {
	useCase usecase.AuthUseCase
}

func NewHandler(uc usecase.AuthUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmid := middleware.Authenticate(app)
	limiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/auth/health", h.health)

	router.HandlerFunc(http.MethodPost, "/api/v1/auth/register", limiter.Limit(Register(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/login", limiter.Limit(Login(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/refresh", RefreshToken(app))
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/logout", LogoutUser(app))
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/logout-all", authmid(LogoutAllSessions(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/verify-otp", limiter.Limit(VerifyOTPHandler(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/request-otp", limiter.Limit(RequestOTPHandler(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
