package farms

import (
	"naevis/infra"
	"naevis/internal/farms/usecase"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewFarmUseCase()).registerRoutes(router, app)
}
