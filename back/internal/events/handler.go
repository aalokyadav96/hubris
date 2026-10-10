package events

import (
	"net/http"

	"naevis/infra"
)

func GetEvents(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"events":[]}`))
	}
}

func GetEventsCount(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"count":0}`))
	}
}

func CreateEvent(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"created"}`))
	}
}

func GetEvent(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"event":null}`))
	}
}

func EditEvent(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"updated"}`))
	}
}

func DeleteEvent(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"deleted"}`))
	}
}
