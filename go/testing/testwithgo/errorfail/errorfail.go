// Package errorfail demonstrates appropriate scenarios for t.Fatal() and t.Error()
package errorfail

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		FirstName: "Mike",
		LastName: "Donnici",
		Age: 47,
	}
	json.NewEncoder(w).Encode(p)
}
