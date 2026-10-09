package routes

import (
	"context"
	"net/http"
	"time"

	"naevis/infra"
	"naevis/middleware"

	"github.com/julienschmidt/httprouter"
)

func SetupRouter(router *httprouter.Router, app *infra.Deps) {
	// Liveness
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// Readiness: check Redis and Postgres (if configured)
	router.GET("/ready", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if app.Cache != nil {
			if _, err := app.Cache.Ping(ctx); err != nil {
				http.Error(w, "redis unavailable", http.StatusServiceUnavailable)
				return
			}
		}

		if app.DB != nil {
			if err := app.DB.Ping(ctx); err != nil {
				http.Error(w, "db unavailable", http.StatusServiceUnavailable)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	// Attach structured middleware to all place endpoints.
	RegisterRoutes(router, app)

	// Example of attaching additional global middleware to the router itself.
	router.NotFound = middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
}
