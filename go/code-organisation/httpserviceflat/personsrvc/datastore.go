package personsrvc

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

// Datastore contains connections to the data sources required for the service
type Datastore struct {
	MySQL *MySQLConnection
	Mongo *MongoConnection
}

// MySQLConnection holds values for a connection to a MySQL server
type MySQLConnection struct {
	DSN     string // Data source name, format: user:pass@tcp(hostname:3306)/ - exclude database name
	DBName  string // Database name
	Session *sql.DB
}

// NewMySQLConnection returns a pointer to a MySQL Connection
func NewMySQLConnection(dsn, dbName string) (*MySQLConnection, error) {
	m := &MySQLConnection{
		DSN:    dsn,
		DBName: dbName,
	}
	err := m.checkFields()
	if err != nil {
		return m, err
	}

	// sql.Open() doesn't actually verify a connection, so add a Ping()
	// to verify the connection has been made.
	m.Session, err = sql.Open("mysql", m.DSN+m.DBName)
	if err != nil {
		return m, err
	}
	err = m.Session.Ping()

	return m, err
}

func (m *MySQLConnection) checkFields() error {
	if m.DSN == "" {
		return errors.New("MySQL DSN (store source name / connection string) not set")
	}
	if m.DBName == "" {
		return errors.New("MySQL db name not set")
	}
	return nil
}

// MongoConnection represents a connection to a Mongo server
// and includes convenience methods for accessing each collection
type MongoConnection struct {
	DSN     string
	DBName  string
	Session *mgo.Session
}

// NewMongoConnection returns a pointer to a MongoDB Connection
func NewMongoConnection(dsn, dbname string) (*MongoConnection, error) {

	m := &MongoConnection{
		DSN:    dsn,
		DBName: dbname,
	}

	err := m.checkFields()
	if err != nil {
		return m, err
	}

	// mgo.Dial() may not actually verify the connection so add a Ping()
	m.Session, err = mgo.Dial(m.DSN)
	if err != nil {
		return m, err
	}
	err = m.Session.Ping()

	return m, err
}

// Collection returns a pointer to the collection specified by name
func (m *MongoConnection) Collection(name string) (*mgo.Collection, error) {
	return m.Session.DB(m.DBName).C(name), nil
}

func (m *MongoConnection) checkFields() error {
	if m.DSN == "" {
		return errors.New("MongoDB DSN (data source name / connection string) not set")
	}
	if m.DBName == "" {
		return errors.New("MongoDB db name not set")
	}

	return nil
}
