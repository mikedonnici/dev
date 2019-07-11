package personsrvc

import (
	"gopkg.in/mgo.v2/bson"
)

// Person holds store for a person
type Person struct {
	ID        int           `json:"id" bson:"id"`
	IOD       bson.ObjectId `json:"oid,omitempty" bson:"_id,omitempty"`
	FirstName string        `json:"firstName" bson:"firstname"`
	LastName  string        `json:"lastName" bson:"lastname"`
	Age       int           `json:"age" bson:"age"`
}

// PersonByID fetches a person record from MySQL
func (d *Datastore) PersonByID(id string) (Person, error) {
	var p Person
	q := "SELECT id, firstname, lastname, age FROM people WHERE id = ?"
	err := d.MySQL.Session.QueryRow(q, id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.Age)
	return p, err
}

// PersonByOID fetches a person record from MongoDB
func (d *Datastore) PersonByOID(oid string) (Person, error) {
	var p Person
	q := bson.M{"_id": bson.ObjectIdHex(oid)}
	err := d.Mongo.Session.DB(d.Mongo.DBName).C("people").Find(q).One(&p)
	return p, err
}

// People fetches all the people
func (d *Datastore) People() ([]Person, error) {
	var xp []Person
	q := bson.M{}
	err := d.Mongo.Session.DB(d.Mongo.DBName).C("people").Find(q).All(&xp)
	return xp, err
}
