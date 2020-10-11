package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhanusaputra/anywhat-server/pkg/env"
	"github.com/dhanusaputra/anywhat-server/util/testutil"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestAddAuth(t *testing.T) {
	tests := []struct {
		name string
		req  func() *http.Request
		mock func()
		want string
	}{
		{
			name: "happy path",
			req: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjpudWxsLCJpZCI6IjEiLCJpc3MiOiJhbnl3aGF0IiwidXNlcm5hbWUiOiJ1c2VybmFtZSJ9.6rNiOIRQ_s5hrDDFw0QdRtmHpz8DC_w22oNRVJzxOiY")
				return req
			},
			mock: func() {
				env.AuthEnable = true
			},
			want: "",
		},
		{
			name: "auth disable",
			req: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				return req
			},
			mock: func() {
				env.AuthEnable = false
			},
			want: "",
		},
		{
			name: "empty header",
			req: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				return req
			},
			mock: func() {
				env.AuthEnable = true
			},
			want: "",
		},
		{
			name: "validate failed",
			req: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjpudWxsLCJpZCI6IjEiLCJpc3MiOiJhbnl3aGF0IiwidXNlcm5hbWUiOiJ1c2VybmFtZSJ9.6rNiOIRQ_s5hrDDFw0QdRtmHpz8DC_w22oNRVJzxOi")
				return req
			},
			mock: func() {
				env.AuthEnable = true
			},
			want: "signature is invalid",
		},
	}

	for _, tt := range tests {
		defer testutil.NewPtrs([]interface{}{&env.AuthEnable}).Restore()
		tt.mock()

		w := httptest.NewRecorder()

		r := chi.NewRouter()

		r.Use(AddAuth)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
		r.ServeHTTP(w, tt.req())

		got := w.Body.String()
		assert.Equal(t, tt.want, got, tt.name)
	}
}
