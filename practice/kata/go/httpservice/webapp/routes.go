package webapp

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

const validToken = "abc123"
const inValidToken = "abc124"

// routes sets up the routes and handlers. The handleXXX functions are invoked and return a HandlerFunc.
func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/hello/{name}", s.handleHelloName())
	s.router.HandleFunc("/hello.json", s.handleHelloJSON())
	s.router.HandleFunc("/adminok", s.adminAuth(s.handleAdmin(), validToken))
	s.router.HandleFunc("/adminfail", s.adminAuth(s.handleAdmin(), inValidToken))
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusFound)
	}
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

func (s *server) handleHelloName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		xs := mux.Vars(r)
		io.WriteString(w, "hello "+xs["name"])
	}
}

func (s *server) handleHelloJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"response": "hello"}`)
	}
}

func (s *server) handleAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "admin authenticated")
	}
}

func (s *server) adminAuth(h http.HandlerFunc, token string) http.HandlerFunc {

	if token == validToken {
		return h
	}

	return func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}
}
