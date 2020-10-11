package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestAddRequestID(t *testing.T) {
	tests := []struct {
		name string
		req  func() *http.Request
	}{
		{
			name: "happy path",
			req: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Add("X-Request-Id", "req-123456")

				return req
			},
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()

		r := chi.NewRouter()

		r.Use(AddRequestID)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			requestID := GetReqID(r.Context())
			response := fmt.Sprintf("RequestID: %s", requestID)

			w.Write([]byte(response))
		})
		r.ServeHTTP(w, tt.req())

		got := w.Body.String()
		assert.NotEmpty(t, got)
	}
}
