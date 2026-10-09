package routes

import (
	"naevis/infra"
	"naevis/internal/places"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, app *infra.Deps) {

	places.RegisterRoutes(router, app)
}
