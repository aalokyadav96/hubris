package profile

import (
	"naevis/infra"
	"naevis/internal/profile/usecase"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewProfileUseCase()).registerRoutes(router, app)
}
