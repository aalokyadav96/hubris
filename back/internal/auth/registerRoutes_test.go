package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"naevis/infra"

	"github.com/julienschmidt/httprouter"
)

func TestRegisterRoutesIncludesAuthEndpoints(t *testing.T) {
	router := httprouter.New()
	RegisterRoutes(router, &infra.Deps{})

	paths := []string{
		"/api/v1/auth/register",
		"/api/v1/auth/login",
		"/api/v1/auth/refresh",
		"/api/v1/auth/logout",
		"/api/v1/auth/logout-all",
		"/api/v1/auth/verify-otp",
		"/api/v1/auth/request-otp",
	}

	for _, path := range paths {
		req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(`{}`))
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code == http.StatusNotFound {
			t.Fatalf("expected route %s to be registered", path)
		}
	}
}
