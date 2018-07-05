package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/server"
)

// TestNewServer creates a new server and tests the root path for a 200 ok response
func TestNewServer(t *testing.T) {
	is := is.New(t)
	s := server.NewServer("8888", nil)
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	is.Equal(w.Code, 200) // expected 200 ok
}
