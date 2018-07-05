package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
)

type server struct {
	port   string
	router *mux.Router
	store  *datastore.Datastore
}

// NewServer returns a pointer to an initialised server with a connected datastore
func NewServer(port string, store *datastore.Datastore) *server {
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
