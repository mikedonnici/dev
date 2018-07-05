package datastore_test

import (
	"log"
	"testing"

	"github.com/matryer/is"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mongo"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mysql"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/testdata"
)

var testDB = testdata.New()
var ds = datastore.New()

func TestPerson(t *testing.T) {

	var err error

	// install databases
	err = setupDatabases()
	if err != nil {
		log.Fatalln(err)
	}
	defer teardownDatabases()

	// connect datastore
	err = datastoreConnectMySQL()
	if err != nil {
		log.Fatalln(err)
	}
	err = datastoreConnectMongoDB()
	if err != nil {
		log.Fatalln(err)
	}

	t.Run("group", func(t *testing.T) {
		t.Run("testPersonByID", testPersonByID)
		t.Run("testPersonByOID", testPersonByOID)
		t.Run("testPeople", testPeople)
	})
}

// setUpDatabases creates and populates test databases
func setupDatabases() error {
	err := testDB.SetupMySQL()
	if err != nil {
		return err
	}
	return testDB.SetupMongoDB()
}

// teardownDatabases cleans up the test databases
func teardownDatabases() {
	testDB.TearDownMySQL()
	testDB.TearDownMongoDB()
}

// datastoreConnectMySQL connects the datastore to the MySQL test database
func datastoreConnectMySQL() error {
	var err error
	ds.MySQL, err = mysql.NewConnection(testdata.MySQLDSN, testDB.DBName, "test mysql db")
	return err
}

// datastoreConnectMongoDB connects the datastore to the test Mongo database
func datastoreConnectMongoDB() error {
	var err error
	ds.Mongo, err = mongo.NewConnection(testdata.MongoDSN, testDB.DBName, "test")
	return err
}

// Test fetch person
func testPersonByID(t *testing.T) {
	is := is.New(t)

	cases := []struct {
		id        string
		firstName string
	}{
		{id: "1", firstName: "Broderick"},
		{id: "5", firstName: "Declan"},
	}

	for _, c := range cases {
		p, err := ds.PersonByID(c.id)
		is.NoErr(err)                      // error fetching person by id
		is.Equal(p.FirstName, c.firstName) // incorrect first name
	}
}

// Test fetch person from MongoDb by OID
func testPersonByOID(t *testing.T) {
	is := is.New(t)

	cases := []struct {
		oid       string
		firstName string
	}{
		{oid: "5b3bcd72463cd6029e04de18", firstName: "Broderick"},
		{oid: "5b3bcd72463cd6029e04de20", firstName: "Declan"},
	}

	for _, c := range cases {
		p, err := ds.PersonByOID(c.oid)
		is.NoErr(err)                      // error fetching person by object id
		is.Equal(p.FirstName, c.firstName) // incorrect first name
	}
}

// Test fetch people
func testPeople(t *testing.T) {
	is := is.New(t)
	xp, err := ds.People()
	is.NoErr(err)        // error fetching people
	is.Equal(len(xp), 5) // expected 5 people records
}
