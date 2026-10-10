package crops

import (
	"naevis/infra"
	"naevis/internal/crops/usecase"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewCropUseCase()).registerRoutes(router, app)
}
