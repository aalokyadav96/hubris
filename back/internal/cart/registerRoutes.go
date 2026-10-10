package cart

import (
	"naevis/infra"
	"naevis/internal/cart/usecase"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewCartUseCase()).registerRoutes(router, app)
}
