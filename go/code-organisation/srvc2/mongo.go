package main

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

// MongoDBConnection represents a connection to a Mongo server
// and includes convenience methods for accessing each collection
type MongoDBConnection struct {
	DSN     string
	DBName  string
	Session *mgo.Session
}

func NewMongoDBConnection(dsn, dbname string) (*MongoDBConnection, error) {

	m := &MongoDBConnection{
		DSN:    dsn,
		DBName: dbname,
	}
	err := m.checkFields()
	if err != nil {
		return m, err
	}

	m.Session, err = mgo.Dial(m.DSN)
	return m, err
}

// Collection returns a pointer to the collection specified by name
func (m *MongoDBConnection) Collection(name string) (*mgo.Collection, error) {
	return m.Session.DB(m.DBName).C(name), nil
}

func (m *MongoDBConnection) checkFields() error {
	if m.DSN == "" {
		return errors.New("MongoDB DSN (data source name / connection string) not set")
	}
	if m.DBName == "" {
		return errors.New("MongoDB db name not set")
	}
	return nil
}
