package webapp

import (
	"testing"
	"github.com/matryer/is"
	"net/http/httptest"
	"strings"
)

func TestNew(t *testing.T) {
	srv := New()
	t.Log(srv)
}

func TestNotFound(t *testing.T) {
	is := is.New(t)
	srv := New()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/notfound", nil)
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 404) // should get 404 Not Found
}

func TestHello(t *testing.T) {
	is := is.New(t)
	srv := New()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/hello", nil)
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 200)
}

func TestIndex(t *testing.T) {
	is := is.New(t)
	srv := New()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 301)
	is.Equal(w.Result().Header.Get("Location"), "/hello")
}

func TestHelloJSON(t *testing.T) {
	is := is.New(t)
	srv := New()
	r := httptest.NewRequest("GET", "/hello.json", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.HeaderMap.Get("Content-Type"), "application/json")
	body := strings.Replace(w.Body.String(), " ", "", -1)
	is.Equal(body, `{"response":"hello"}`)
}

func TestAdminMiddleware(t *testing.T) {
	is := is.New(t)
	srv := New()

	r := httptest.NewRequest("GET", "/admin/abc123", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 200)
	body := w.Body.String()
	is.Equal(body, "admin authorized")

	r = httptest.NewRequest("GET", "/admin/abc124", nil)
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 401)
}


