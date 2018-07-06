package testdata

import (
	"fmt"
	"encoding/json"
	"time"
	"database/sql"

	"github.com/hashicorp/go-uuid"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Hard coded for local dev and Travis CI
const MySQLDSN = "root:password@tcp(localhost:3306)/"
const MongoDSN = "mongodb://localhost"

type TestStore struct {
	DBName         string
	MySQLSession   *sql.DB
	MongoDBSession *mgo.Session
}

// New returns a pointer to a TestStore
func New() *TestStore {

	s, _ := uuid.GenerateUUID()
	n := fmt.Sprintf("%v_test", s[0:7])

	t := TestStore{
		DBName: n,
	}
	return &t
}

// SetupMySQL creates and populates the test MySQL database
func (t *TestStore) SetupMySQL() error {

	var err error

	t.MySQLSession, err = sql.Open("mysql", MySQLDSN)
	time.Sleep(3 * time.Second) // give it a sec to connect to server
	_, err = t.MySQLSession.Exec(fmt.Sprintf(CREATE_MYSQL_DB, t.DBName))
	if err != nil {
		return errors.Wrap(err, "Error creating test schema")
	}

	// connect session to newly create database
	dsn := MySQLDSN + t.DBName
	t.MySQLSession, err = sql.Open("mysql", dsn)
	if err != nil {
		t.TearDownMySQL()
		return errors.Wrap(err, "Error connecting to the test database")
	}

	t.MySQLSession.Exec(fmt.Sprintf(CREATE_MYSQL_TABLE, t.DBName))
	if err != nil {
		t.TearDownMySQL()
		return errors.Wrap(err, "Error creating tables")
	}

	_, err = t.MySQLSession.Exec(fmt.Sprintf(INSERT_MYSQL_DATA, t.DBName))
	if err != nil {
		t.TearDownMySQL()
		return errors.Wrap(err, "Error inserting data")
	}

	// Need a little delay to ensure the database is ready
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Duration(time.Second))
	}

	return nil
}

func (t *TestStore) TearDownMySQL() error {
	_, err := t.MySQLSession.Exec(fmt.Sprintf(DROP_MYSQL_DB, t.DBName))
	if err != nil {
		return errors.Wrap(err, "Error deleting MySQL test database")
	}
	return nil
}

// SetupMongoDB connects to the test database and populates a collection
func (t *TestStore) SetupMongoDB() error {

	var err error

	t.MongoDBSession, err = mgo.Dial(MongoDSN)
	if err != nil {
		return errors.Wrap(err, "Error establishing session with Mongo")
	}

	err = t.MongoDBSession.Ping()
	if err != nil {
		return errors.Wrap(err, "Error pinging Mongo")
	}

	var xp []struct {
		OID       bson.ObjectId `json:"oid" bson:"_id"`
		ID        int           `json:"id" bson:"id"`
		FirstName string        `json:"firstname" bson:"firstname"`
		LastName  string        `json:"lastname" bson:"lastname"`
		Age       int           `json:"age" bson:"age"`
	}
	err = json.Unmarshal([]byte(MONGO_DATA), &xp)
	if err != nil {
		return errors.Wrap(err, "Unmarshal error")
	}

	for _, p := range xp {
		err = t.MongoDBSession.DB(t.DBName).C(MONGO_COLLECTION).Insert(p)
		if err != nil {
			return errors.Wrap(err, "Error inserting test docs into mongo")
		}
	}

	return nil
}

func (t *TestStore) TearDownMongoDB() error {
	err := t.MongoDBSession.DB(t.DBName).DropDatabase()
	if err != nil {
		return errors.Wrap(err, "Error deleting Mongo test database")
	}
	return nil
}
