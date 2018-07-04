package datastore

import "gopkg.in/mgo.v2/bson"

// Person holds data for a person
type Person struct {
	ID        int    `json:"id" bson:"id"`
	IOD       string `json:"oid" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstname"`
	LastName  string `json:"lastName" bson:"lastname"`
	Age       int    `json:"age" bson:"age"`
}

// PersonByID fetches a person record from MySQL
func (d *Datastore) PersonByID(id int) (Person, error) {
	var p Person
	q := "SELECT id, firstname, lastname, age FROM people WHERE id = ?"
	err := d.MySQL.Session.QueryRow(q, id).Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.Age,
	)
	return p, err
}

// PersonByOID fetches a person record from MongoDB
func (d *Datastore) PersonByOID(oid string) (Person, error) {
	var p Person
	q := bson.M{"_id": bson.ObjectIdHex(oid)}
	err := d.Mongo.Session.DB(d.Mongo.DBName).C("people").Find(q).One(&p)
	return p, err
}
