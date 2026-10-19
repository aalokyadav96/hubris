package routes

import (
	"naevis/infra"
	"naevis/internal/auth"
	"naevis/internal/cart"
	"naevis/internal/crops"
	"naevis/internal/farms"
	"naevis/internal/profile"
	"naevis/internal/tickets"
	"naevis/internal/timetable"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {
	auth.RegisterRoutes(router, app)
	farms.RegisterRoutes(router, app)
	crops.RegisterRoutes(router, app)
	timetable.RegisterRoutes(router, app)
	tickets.RegisterRoutes(router, app)
	cart.RegisterRoutes(router, app)
	profile.RegisterRoutes(router, app)
}
