package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNormalizeTrailingSlashes(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{path: "/", want: "/"},
		{path: "//", want: "/"},
		{path: "///", want: "/"},
		{path: "/mcp", want: "/mcp"},
		{path: "/mcp/", want: "/mcp"},
		{path: "/mcp//", want: "/mcp"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.want {
					t.Errorf("handler path = %q, want %q", r.URL.Path, tt.want)
				}
				w.WriteHeader(http.StatusNoContent)
			})

			recorder := httptest.NewRecorder()
			normalizeTrailingSlashes(mux).ServeHTTP(recorder, httptest.NewRequest(http.MethodPost, "http://example.com"+tt.path, nil))

			if recorder.Code != http.StatusNoContent {
				t.Errorf("status = %d, want %d", recorder.Code, http.StatusNoContent)
			}
			if location := recorder.Header().Get("Location"); location != "" {
				t.Errorf("unexpected redirect to %q", location)
			}
		})
	}
}
