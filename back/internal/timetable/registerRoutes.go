package timetable

import (
	"naevis/infra"
	"naevis/internal/timetable/usecase"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	NewHandler(usecase.NewTimetableUseCase()).registerRoutes(router, app)
}
