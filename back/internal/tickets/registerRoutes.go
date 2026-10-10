package tickets

import (
	"naevis/infra"
	"naevis/internal/tickets/usecase"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewTicketUseCase()).registerRoutes(router, app)
}
