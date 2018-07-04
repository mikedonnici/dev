package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
)

type server struct {
	store  *datastore.Datastore
	router *mux.Router
}

// Start fires up the http server on the specified port and associates it with the specified datastore.
func Start(port string, store *datastore.Datastore) error {
	s := server{
		store: store,
		router: mux.NewRouter(),
	}
	s.routes()
	return http.ListenAndServe(":" + port, s.router)
}
