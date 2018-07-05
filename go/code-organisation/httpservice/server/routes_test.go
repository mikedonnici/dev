package server_test

import (
	"testing"
	"log"
	"strings"
	"net/http/httptest"

	"github.com/mikedonnici/dev/go/code-organisation/httpservice/testdata"
	"github.com/matryer/is"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/server"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
)

var testDB = testdata.New()
var ds = datastore.New()

// TestRoutes sets up test databases, connects a testDB to the database and starts a server with the datastore.
// It then runs a group of route tests and tears down the test databases.
func TestRoutes(t *testing.T) {

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

	// run tests
	t.Run("group", func(t *testing.T) {
		t.Run("groupTestPersonByID", groupTestPersonByID)
		t.Run("groupTestPeople", groupTestPeople)
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
	ds.MySQL.DSN = testdata.MySQLDSN
	ds.MySQL.DBName = testDB.DBName
	ds.MySQL.Desc = "test"
	return ds.MySQL.Connect()
}

// datastoreConnectMongoDB connects the datastore to the test Mongo database
func datastoreConnectMongoDB() error {
	ds.Mongo.DSN = testdata.MongoDSN
	ds.Mongo.DBName = testDB.DBName
	ds.Mongo.Desc = "test"
	return ds.Mongo.Connect()
}

func groupTestPersonByID(t *testing.T) {
	is := is.New(t)
	r := httptest.NewRequest("GET", "/person/1", nil)
	w := httptest.NewRecorder()
	srv := server.NewServer("8888", ds)
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 200) // response not 200 ok
	s := `{"id":1,"firstName":"Broderick","lastName":"Reynolds","age":68}`
	expect := stringRemove(s, " ", "\n")
	body := stringRemove(w.Body.String(), " ", "\n")
	is.Equal(body, expect) // response not as expected
}

func groupTestPeople(t *testing.T) {
	is := is.New(t)
	r := httptest.NewRequest("GET", "/people", nil)
	w := httptest.NewRecorder()
	srv := server.NewServer("8888", ds)
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, 200) // response not 200 ok
}

// stringRemove removes one or more character from a string - for cleaning up and comparing reponses.
func stringRemove(str string, removeChars ...string) string {
	for _, c := range removeChars {
		str = strings.Replace(str, c, "", -1)
	}
	return str
}
