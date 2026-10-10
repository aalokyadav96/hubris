package auth

import (
	"net/http"

	"naevis/infra"
)

func stubAuthHandler(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, msg, http.StatusNotImplemented)
	}
}

func Register(app *infra.Deps) http.HandlerFunc {
	_ = app
	return stubAuthHandler("auth register not implemented")
}

func Login(app *infra.Deps) http.HandlerFunc {
	_ = app
	return stubAuthHandler("auth login not implemented")
}

func RefreshToken(app *infra.Deps) http.HandlerFunc {
	_ = app
	return stubAuthHandler("auth refresh not implemented")
}

func LogoutUser(app *infra.Deps) http.HandlerFunc {
	_ = app
	return stubAuthHandler("auth logout not implemented")
}

func LogoutAllSessions(app *infra.Deps) http.HandlerFunc {
	_ = app
	return stubAuthHandler("auth logout-all not implemented")
}

func VerifyOTPHandler(app *infra.Deps) http.HandlerFunc {
	_ = app
	return stubAuthHandler("auth verify-otp not implemented")
}

func RequestOTPHandler(app *infra.Deps) http.HandlerFunc {
	_ = app
	return stubAuthHandler("auth request-otp not implemented")
}
