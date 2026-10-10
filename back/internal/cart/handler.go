package cart

import (
	"net/http"

	"naevis/infra"
)

func GetCartItems(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"items":[]}`))
	}
}

func GetCartItemsCount(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"count":0}`))
	}
}

func CreateCartItem(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"created"}`))
	}
}

func GetCartItem(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"item":null}`))
	}
}

func EditCartItem(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"updated"}`))
	}
}

func DeleteCartItem(app *infra.Deps) http.HandlerFunc {
	_ = app
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"deleted"}`))
	}
}
