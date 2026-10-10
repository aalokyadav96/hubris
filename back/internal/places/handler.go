package places

import (
	"net/http"

	"naevis/infra"
)

func GetPlaces(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"places":[]}`))
	}
}

func GetPlacesCount(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"count":0}`))
	}
}

func CreatePlace(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"created"}`))
	}
}

func GetPlace(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"place":null}`))
	}
}

func EditPlace(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"updated"}`))
	}
}

func DeletePlace(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"deleted"}`))
	}
}
