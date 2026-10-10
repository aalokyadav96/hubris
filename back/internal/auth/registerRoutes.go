package auth

import (
	"naevis/infra"
	"naevis/internal/auth/usecase"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewAuthUseCase()).registerRoutes(router, app)
}
