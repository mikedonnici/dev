package personsrvc

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	port   string
	router *mux.Router
	store  Datastore
}

// NewServer returns a pointer to an initialised server with a connected datastore
func NewServer(port string, store Datastore) *server {
	s := &server{
		port:   port,
		store:  store,
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

// Start fires up the http server
func (s *server) Start() error {
	return http.ListenAndServe(":"+s.port, s.router)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
