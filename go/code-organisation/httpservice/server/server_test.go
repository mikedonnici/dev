package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/mikedonnici/dev/go/code-organisation/httpservice/server"
)

// TestNewServer creates a new server and tests the root path for a 200 ok response
func TestNewServer(t *testing.T) {
	s := server.NewServer("8888", nil)
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	want := 200
	got := w.Code
	if got != want {
		t.Fatalf("Response Code = %d, want %d", got, want)
	}
}
