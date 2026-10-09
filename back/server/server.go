package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"naevis/infra"
	"naevis/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// Run starts the HTTP server and handles graceful shutdown on SIGINT/SIGTERM.
func Run(deps *infra.Deps) error {
	if deps == nil {
		return nil
	}

	router := httprouter.New()
	routes.SetupRouter(router, deps)

	handler := cors.Default().Handler(router)

	srv := &http.Server{
		Addr:    deps.Config.PORT,
		Handler: handler,
	}

	errc := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errc <- err
		}
		close(errc)
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-sig:
		if deps.Log != nil {
			deps.Log.Sugar().Infow("signal received", "signal", s)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
		_ = deps.Close(ctx)
		if deps.Log != nil {
			_ = deps.Log.Sync()
		}
		return nil
	case err := <-errc:
		return err
	}
}
