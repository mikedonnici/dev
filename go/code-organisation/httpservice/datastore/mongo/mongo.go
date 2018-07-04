package mongo

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

// MongoDBConnection represents a connection to a Mongo server
// and includes convenience methods for accessing each collection
type Connection struct {
	DSN     string
	DBName  string
	Desc    string
	Session *mgo.Session
}

func NewConnection(dsn, dbname, desc string) (Connection, error) {

	m := Connection{
		DSN: dsn,
		DBName: dbname,
		Desc: desc,
	}
	err := m.checkFields()
	if err != nil {
		return m, err
	}

	m.Session, err = mgo.Dial(m.DSN)
	return m, err
}

// Connect to Mongo
func (m *Connection) Connect() error {
	err := m.checkFields()
	if err != nil {
		return err
	}
	m.Session, err = mgo.Dial(m.DSN)
	return err
}

// Collection returns a pointer to the collection specified by name
func (m *Connection) Collection(name string) (*mgo.Collection, error) {
	return m.Session.DB(m.DBName).C(name), nil
}

// Close terminates the Session
//func (m *MongoDBConnection) Close() {
//	m.Session.Close()
//}

func (m *Connection) checkFields() error {
	if m.DSN == "" {
		return errors.New("MongoDBConnection.DSN (data source name / connection string) is not set")
	}
	if m.DBName == "" {
		return errors.New("MongoDBConnection.DBName is not set")
	}
	if m.Desc == "" {
		return errors.New("MongoDBConnection.Desc is not set")
	}
	return nil
}
