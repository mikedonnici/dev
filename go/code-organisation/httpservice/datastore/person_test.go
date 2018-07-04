package datastore_test

import (
	"testing"
	"log"

	"github.com/matryer/is"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/testdata"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
)

// Test fetch person
func TestPersonByID(t *testing.T) {
	is := is.New(t)

	// create test MySQL database
	data := testdata.New()
	err := data.SetupMySQL()
	if err != nil {
		log.Fatalln(err)
	}
	defer data.TearDownMySQL()

	// create a datastore connected to the test database
	ds := datastore.New()
	ds.MySQL.DSN = testdata.MySQLDSN
	ds.MySQL.DBName = data.DBName
	ds.MySQL.Desc = "test"
	err = ds.MySQL.Connect()
	is.NoErr(err) // could not connect to test MySQL database

	cases := []struct {
		id        int
		firstName string
	}{
		{id: 1, firstName: "Broderick"},
		{id: 5, firstName: "Declan"},
	}

	for _, c := range cases {
		p, err := ds.PersonByID(c.id)
		is.NoErr(err)                      // error fetching person by id
		is.Equal(p.FirstName, c.firstName) // incorrect first name
	}
}

// Test fetch person from MongoDb by OID
func TestPersonByOID(t *testing.T) {
	is := is.New(t)

	// create test MySQL database
	data := testdata.New()
	err := data.SetupMongoDB()
	if err != nil {
		log.Fatalln(err)
	}
	//defer data.TearDownMongoDB()

	// create a datastore connected to the test database
	ds := datastore.New()
	ds.Mongo.DSN = testdata.MongoDSN
	ds.Mongo.DBName = data.DBName
	ds.Mongo.Desc = "test"
	err = ds.Mongo.Connect()
	is.NoErr(err) // could not connect to test MySQL database

	cases := []struct {
		oid       string
		firstName string
	}{
		{oid: "5b3bcd72463cd6029e04de18", firstName: "Broderick"},
		{oid: "5b3bcd72463cd6029e04de20", firstName: "Declan"},
	}

	for _, c := range cases {
		p, err := ds.PersonByOID(c.oid)
		is.NoErr(err)                      // error fetching person by id
		is.Equal(p.FirstName, c.firstName) // incorrect first name
	}
}
