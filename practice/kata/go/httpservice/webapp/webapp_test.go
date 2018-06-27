package webapp

import (
	"testing"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"strings"
)

// TestNew tests the creation of a server value
func TestNew(t *testing.T) {
	is := is.New(t)
	srv := New("8080")
	is.Equal(srv.port, "8080")
}

// TestIndexHandler tests the index handler which should return a 302 status and redirect
func TestIndexHandler(t *testing.T) {
	is := is.New(t)
	srv := New("8080")
	r, err := http.NewRequest("GET", "/", nil)
	is.NoErr(err) // request error
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 302) // status code should be 302
}

func TestHelloHandler(t *testing.T) {
	is := is.New(t)
	srv := New("8080")
	r, err := http.NewRequest("GET", "/hello", nil)
	is.NoErr(err) // request error
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 200) // status not 200 OK
	body := w.Body.String()
	is.Equal(body, "hello") // body should be "hello"
}

func TestHelloMikeHandler(t *testing.T) {
	is := is.New(t)
	srv := New("")
	r, err := http.NewRequest("GET", "/hello/mike", nil)
	is.NoErr(err) // request error
	w := httptest.NewRecorder()
	srv.ServeHTTP(w,r)
	is.Equal(w.Code, 200) // status not 200 OK
}

func TestHelloJSONHandler(t *testing.T) {
	is := is.New(t)
	srv := New("")
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/hello.json", nil)
	is.NoErr(err) // request error
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 200) // status should be 200 OK
	is.Equal(w.Header().Get("Content-Type"), "application/json") // wrong content type header
	body := strings.Replace(w.Body.String(), " ", "", -1)
	is.Equal(body, `{"response":"hello"}`)
}

func TestAdminMiddleware(t *testing.T) {
	is := is.New(t)
	srv := New("")

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/admin/abc123", nil)
	is.NoErr(err) // request error
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 200)

	w = httptest.NewRecorder()
	r, err = http.NewRequest("GET", "/admin/abc124", nil)
	is.NoErr(err) // request error
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 401) // expected 401 unauthorized
}

