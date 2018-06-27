package webapp

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type server struct {
	router *mux.Router
}

func Start() {
	s := server{}
	s.router = mux.NewRouter()
	s.routes()
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
