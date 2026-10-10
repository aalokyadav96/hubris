package tickets

import (
	"net/http"

	"naevis/infra"
)

func GetTickets(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"tickets":[]}`))
	}
}

func GetTicketsCount(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"count":0}`))
	}
}

func CreateTicket(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"created"}`))
	}
}

func GetTicket(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ticket":null}`))
	}
}

func EditTicket(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"updated"}`))
	}
}

func DeleteTicket(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"deleted"}`))
	}
}
