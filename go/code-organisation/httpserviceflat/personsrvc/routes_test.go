package personsrvc_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mikedonnici/dev/go/code-organisation/httpserviceflat/personsrvc"
	"github.com/mikedonnici/dev/go/code-organisation/httpserviceflat/personsrvc/testdata"
)

var routesStore personsrvc.Datastore

// TestRoutes sets up test databases, connects a testDB to the database and starts a server with the datastore.
// It then runs a group of route tests and tears down the test databases.
func TestRoutes(t *testing.T) {

	var teardown func()
	routesStore, teardown = testdata.Setup(t)
	defer teardown()

	t.Run("routes", func(t *testing.T) {
		t.Run("routesPersonByID", testRoutesPersonByID)
		t.Run("routesPeople", testRoutesPeople)
	})
}

func testRoutesPersonByID(t *testing.T) {
	r := httptest.NewRequest("GET", "/person/1", nil)
	w := httptest.NewRecorder()
	srv := personsrvc.NewServer("8888", routesStore)
	srv.ServeHTTP(w, r)

	// check reponse code
	want := 200
	got := w.Code
	if got != want {
		t.Fatalf("Response Code = %d, want %d", got, want)
	}

	// check response body
	s := `{"id":1,"firstName":"Broderick","lastName":"Reynolds","age":68}`
	wantBody := stringRemove(s, " ", "\n")
	gotBody := stringRemove(w.Body.String(), " ", "\n")
	if gotBody != wantBody {
		t.Errorf("Response body = %s, want %s", gotBody, wantBody)
	}
}

func testRoutesPeople(t *testing.T) {
	r := httptest.NewRequest("GET", "/people", nil)
	w := httptest.NewRecorder()
	srv := personsrvc.NewServer("8888", routesStore)
	srv.ServeHTTP(w, r)

	want := 200
	got := w.Code
	if got != want {
		t.Fatalf("Response Code = %d, want %d", got, want)
	}
}

// stringRemove removes one or more character from a string and is used to clean
// up response body strings for comparison.
func stringRemove(str string, removeChars ...string) string {
	for _, c := range removeChars {
		str = strings.Replace(str, c, "", -1)
	}
	return str
}
