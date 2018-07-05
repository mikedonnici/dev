package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/arg", s.handleArg("'HELOOOOOOOOOO!!'"))
	s.router.HandleFunc("/person/{id}", s.handlePersonID())
	s.router.HandleFunc("/person/oid/{oid}", s.handlePersonOID())
	s.router.HandleFunc("/people", s.handlePeople())
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

func (s *server) handlePersonID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		person, err := s.store.PersonByID(id)
		respondJSON(w, person, err)
	}
}

func (s *server) handlePersonOID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oid := mux.Vars(r)["oid"]
		if !bson.IsObjectIdHex(oid) {
			respondJSON(w, nil, errors.New("invalid object id"))
			return
		}
		person, err := s.store.PersonByOID(oid)
		respondJSON(w, person, err)
	}
}

func (s *server) handlePeople() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
			h(w, r)
			return
		}
		http.NotFound(w, r)
	}
}

func respondJSON(w http.ResponseWriter, data interface{}, err error) {
	w.Header().Set("content-type", "application/json")
	if err != nil {
		//w.WriteHeader(http.StatusNotFound)
		body := fmt.Sprintf(`{"error": "%s"}`, err.Error())
		io.WriteString(w, body)
		return
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		io.WriteString(w, err.Error())
	}
}
