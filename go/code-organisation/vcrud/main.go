package main

import (
	"fmt"
	"log"
)

// DataManager defines VCRUD behaviour for anything that can manage data
type DataManager interface {
	validate()
	create()
	read()
	update()
	delete()
}

type Datastore struct {
	name string
}

// person represents just the person
type person struct {
	id    string
	name  string
	age   int
	saved bool
}

// PeopleManager does all the grunt work
type PeopleManager struct {
	People []*person
	store  Datastore
}

func (pm *PeopleManager) validate() {
	log.Println("validate() checks the data specific to the receiver value")
}
func (pm *PeopleManager) create() {
	log.Println("create() inserts receiver value data in a store that is attached to the receiver type")
	for _, p := range pm.People {
		log.Printf("Saved %s", p.name)
		p.saved = true
	}
}
func (pm *PeopleManager) read() {
	log.Println("read() queries data from the store that is attached to the receiver type")
}
func (pm *PeopleManager) update() {
	log.Println("update() updates receiver value data in a store that is attached to the receiver type")
}
func (pm *PeopleManager) delete() {
	log.Println("update() deletes receiver data from a store that is attached to the receiver type")
}

func main() {

	ds := Datastore{
		name: "theDB",
	}

	// just a person
	p1 := person{
		id:   "mpd-1970",
		name: "Mike",
		age:  49,
	}
	p2 := person{
		id:   "caw-1975",
		name: "Christie",
		age:  45,
	}
	p3 := person{
		id:   "mgd-2010",
		name: "Maia",
		age:  10,
	}
	p4 := person{
		id:   "lad-2012",
		name: "Leo",
		age:  7,
	}

	m := PeopleManager{
		People: []*person{&p1, &p2, &p3, &p4},
		store:  ds,
	}

	printPeople(m.People)
	saveAll(&m)
	printPeople(m.People)
}

func printPeople(xp []*person) {
	for _, p := range xp {
		fmt.Printf("%s is %v years old and sync'd to datastore: %v\n", p.name, p.age, p.saved)
	}
}

func saveAll(dm DataManager) error {
	dm.create()
	return nil
}
