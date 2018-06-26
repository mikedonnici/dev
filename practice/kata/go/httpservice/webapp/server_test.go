package webapp

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

// TestCreateNewServer creates a server, adds a route to the server
// and then gets the name of the route.
func TestCreateServer(t *testing.T) {
	s := server{}
	t.Log(s)
}

func TestHandleHello(t *testing.T) {
	is := is.New(t)
	s := New("8080")
	r, err := http.NewRequest("GET", "/hello", nil)
	is.NoErr(err) // Request error
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusOK) // Status should be 200 OK
	body := w.Body.String()
	is.Equal(body, "hello") // expected "hello" in the response body
}

func TestHelloJSON(t *testing.T) {
	is := is.New(t)
	s := New("8080")
	r, err := http.NewRequest("GET", "/hello.json", nil)
	is.NoErr(err) // request error
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusOK) // response code
	is.Equal(w.Header().Get("Content-Type"), "application/json") // content type
	body := strings.Replace(w.Body.String(), " ", "", -1)
	is.Equal(body, `{"response":"hello"}`) // response body
}
