package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func TestAuthentication(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := newRouter("X-Test-Token", "expected", zerolog.New(io.Discard))
	for _, test := range []struct {
		name  string
		path  string
		value string
		want  int
	}{
		{"matching header", "/check", "expected", http.StatusOK},
		{"wrong header", "/check", "wrong", http.StatusForbidden},
		{"missing header", "/check", "", http.StatusForbidden},
		{"health without authentication", "/health", "", http.StatusOK},
	} {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, test.path, nil)
			if test.value != "" {
				request.Header.Set("X-Test-Token", test.value)
			}
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			if response.Code != test.want {
				t.Fatalf("status = %d, want %d; body: %s", response.Code, test.want, response.Body.String())
			}
		})
	}
}
