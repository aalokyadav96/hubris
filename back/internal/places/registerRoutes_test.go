package places

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"naevis/infra"

	"github.com/julienschmidt/httprouter"
)

func TestRegisterRoutesIncludesPlaceEndpoints(t *testing.T) {
	router := httprouter.New()
	RegisterRoutes(router, &infra.Deps{})

	tests := []struct {
		method string
		path   string
	}{
		{method: http.MethodGet, path: "/api/v1/places/places"},
		{method: http.MethodGet, path: "/api/v1/places/places/count"},
		{method: http.MethodPost, path: "/api/v1/places/place"},
		{method: http.MethodGet, path: "/api/v1/places/place/123"},
		{method: http.MethodPut, path: "/api/v1/places/place/123"},
		{method: http.MethodDelete, path: "/api/v1/places/place/123"},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(tc.method, tc.path, strings.NewReader(`{}`))
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code == http.StatusNotFound {
			t.Fatalf("expected route %s %s to be registered", tc.method, tc.path)
		}
	}
}
