package datastore_test

import (
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/datastoretest"
	"testing"

	"github.com/matryer/is"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
)

var personStore *datastore.Datastore

func TestPerson(t *testing.T) {

	var teardown func()
	personStore, teardown = datastoretest.Setup(t)
	defer teardown()

	// run tests
	t.Run("group", func(t *testing.T) {
		t.Run("testPersonByID", testPersonByID)
		t.Run("testPersonByOID", testPersonByOID)
		t.Run("testPeople", testPeople)
	})
}

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
		p, err := personStore.PersonByID(c.id)
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
		p, err := personStore.PersonByOID(c.oid)
		is.NoErr(err)                      // error fetching person by object id
		is.Equal(p.FirstName, c.firstName) // incorrect first name
	}
}

// Test fetch people
func testPeople(t *testing.T) {
	is := is.New(t)
	xp, err := personStore.People()
	is.NoErr(err)        // error fetching people
	is.Equal(len(xp), 5) // expected 5 people records
}
