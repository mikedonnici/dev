package main

import "fmt"

// Manager describes the  'VCRUD' interface
type Manager interface {
	Validate(entity interface{}) error
	Create(entity interface{}) error
	Read(entity interface{}) error
	Update(entity interface{}) error
	Delete(entity interface{}) error
}

// Store holds database connections and data managers
type Store struct {
	MySQL    *MySQLConnection
	Mongo    *MongoDBConnection
	Managers map[string]Manager
}

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name""`
}

type PersonManager struct{}

func (pm PersonManager) Validate(entity interface{}) error {
	p := entity.(Person)
	fmt.Println(p)
	return nil
}
func (pm PersonManager) Create(person interface{}) error {
	return nil
}
func (pm PersonManager) Read(person interface{}) error {
	return nil
}
func (pm PersonManager) Update(person interface{}) error {
	return nil
}
func (pm PersonManager) Delete(person interface{}) error {
	return nil
}

func main() {

	pm := PersonManager{}

	crs := Store{
		Managers: map[string]Manager{
			"PersonManager": pm,
		},
	}

	crs.Managers["PersonManager"].Validate(Person{ID: "123abc", Name: "Mike"})
}
