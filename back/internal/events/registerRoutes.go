package events

import (
	"naevis/infra"
	"naevis/internal/events/usecase"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewEventUseCase()).registerRoutes(router, app)
}
