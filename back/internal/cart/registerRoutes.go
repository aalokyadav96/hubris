package cart

import (
	"encoding/json"
	"naevis/infra"
	"naevis/internal/cart/usecase"
	"naevis/middleware"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewCartUseCase()).registerRoutes(router, app)
}

// Handler exposes cart HTTP endpoints.
type Handler struct {
	useCase usecase.CartUseCase
}

func NewHandler(uc usecase.CartUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	authmidware := middleware.Authenticate(app)
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 20, 1*time.Minute, 1000)

	router.GET("/cart/health", h.health)
	router.HandlerFunc(http.MethodGet, "/api/v1/cart/items", rateLimiter.Limit(GetCartItems(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/cart/items/count", rateLimiter.Limit(GetCartItemsCount(app)))
	router.HandlerFunc(http.MethodPost, "/api/v1/cart/item", authmidware(CreateCartItem(app)))
	router.HandlerFunc(http.MethodGet, "/api/v1/cart/item/:itemid", GetCartItem(app))
	router.HandlerFunc(http.MethodPut, "/api/v1/cart/item/:itemid", authmidware(EditCartItem(app)))
	router.HandlerFunc(http.MethodDelete, "/api/v1/cart/item/:itemid", authmidware(DeleteCartItem(app)))
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
