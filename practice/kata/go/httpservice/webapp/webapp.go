package webapp

import (
	"net/http"
	"github.com/gorilla/mux"
	"io"
)

type server struct {
	router *mux.Router
}

func New() *server {
	s := &server{
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	s.router.Handle("/", s.handleIndex())
	s.router.Handle("/hello", s.handleHello())
	s.router.Handle("/hello.json", s.handleHelloJSON())
	s.router.Handle("/admin/{token}", s.admin(s.handleAdmin()))
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/hello")
		w.WriteHeader(301)
	}
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

func (s *server) handleHelloJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"response": "hello"}`)
	}
}

func (s *server) admin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		if v["token"] == "abc123" {
			h(w, r)
		}
		unauthorized(w, r)
	}
}

func (s *server) handleAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "admin authorized")
	}
}

func unauthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}



