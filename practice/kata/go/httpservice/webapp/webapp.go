package webapp

import (
	"net/http"
	"github.com/gorilla/mux"
	"io"
)

type server struct {
	port string
	router *mux.Router
}

// New returns a pointer to a server
func New(port string) *server {
	s := &server{
		port: port,
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *server)routes() {
	s.router.Handle("/", s.handleIndex())
	s.router.Handle("/hello", s.handleHello())
	s.router.Handle("/hello/{name}", s.handleHelloName())
	s.router.Handle("/hello.json", s.handleHelloJSON())
	s.router.Handle("/admin/{token}", s.admin(s.handleAdmin()))
}

// ServeHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusFound)
	}
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,"hello")
	}
}

func (s *server) handleHelloName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello " + mux.Vars(r)["name"])
	}
}

func (s *server) handleHelloJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w,`{"response": "hello"}`)
	}
}

func (s *server) handleAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,"admin authorized")
	}
}

func (s *server) admin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := mux.Vars(r)["token"]
		if token == "abc123" {
			h(w,r)
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}
