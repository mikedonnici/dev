package webapp

import (
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	port   string
	router *mux.Router
}

func New(port string) *server {
	s := &server{
		port:   port,
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *server) Start() {
	http.ListenAndServe(":"+s.port, s.router)
}
