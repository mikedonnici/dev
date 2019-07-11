package personsrvc_test

import (
	"net/http/httptest"
	"testing"

	"github.com/mikedonnici/dev/go/code-organisation/httpserviceflat/personsrvc"
)

// TestNewServer creates a new http server with an empty datastore and checks
// the root path for a 200 ok response.
func TestNewServer(t *testing.T) {
	s := personsrvc.NewServer("8888", personsrvc.Datastore{})
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	want := 200
	got := w.Code
	if got != want {
		t.Fatalf("Response Code = %d, want %d", got, want)
	}
}
