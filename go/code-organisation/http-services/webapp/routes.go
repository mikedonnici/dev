package webapp

import (
	"io"
	"net/http"
)

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/arg", s.handleArg("'HELOOOOOOOOOO!!'"))
	s.router.HandleFunc("/mwtrue", s.middleWare(s.handleMiddleware(), true))
	s.router.HandleFunc("/mwfalse", s.middleWare(s.handleMiddleware(), false))
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "index")
	}
}

func (s *server) handleArg(msg string) http.HandlerFunc {
	msg = "This arg was passed in: " + msg + " and can be accessed in the returned func because we have a closure!"
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, msg)
	}
}

func (s *server) handleMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "middleware ran")
	}
}

// Middleware takes a HandlerFunc arg and returns a HandlerFunc.
// The HandlerFunc that is returned contains the middleware logic
func (s *server) middleWare(h http.HandlerFunc, allow bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if allow {
			h(w,r)
			return
		}
		http.NotFound(w, r)
	}
}