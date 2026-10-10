package cart

import (
	"encoding/json"
	"net/http"

	"naevis/infra"
	"naevis/internal/cart/usecase"

	"github.com/julienschmidt/httprouter"
)

// Handler exposes cart HTTP endpoints.
type Handler struct {
	useCase usecase.CartUseCase
}

func NewHandler(uc usecase.CartUseCase) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) registerRoutes(router *httprouter.Router, app *infra.Deps) {
	_ = app
	router.GET("/cart/health", h.health)
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status, err := h.useCase.Health(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}
