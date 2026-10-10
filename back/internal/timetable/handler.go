package timetable

import (
	"net/http"

	"naevis/infra"
)

func GetTimetables(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"timetables":[]}`))
	}
}

func GetTimetablesCount(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"count":0}`))
	}
}

func CreateTimetable(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"created"}`))
	}
}

func GetTimetable(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"timetable":null}`))
	}
}

func EditTimetable(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"updated"}`))
	}
}

func DeleteTimetable(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"deleted"}`))
	}
}
