package datastore_test

import (
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/datastoretest"
	"testing"

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

	cases := []struct {
		id        string
		firstName string
	}{
		{id: "1", firstName: "Broderick"},
		{id: "5", firstName: "Declan"},
	}

	for _, c := range cases {
		arg := c.id
		want := c.firstName
		p, err := personStore.PersonByID(c.id)
		if err != nil {
			t.Errorf("PersonByID(%s) err = %s", arg, err)
			continue
		}
		got := p.FirstName
		if got != want {
			t.Errorf("PersonByID(%s).FirstName = %q, want %q", arg, got, want)
		}
	}
}

func testPersonByOID(t *testing.T) {

	cases := []struct {
		oid       string
		firstName string
	}{
		{oid: "5b3bcd72463cd6029e04de18", firstName: "Broderick"},
		{oid: "5b3bcd72463cd6029e04de20", firstName: "Declan"},
	}

	for _, c := range cases {
		arg := c.oid
		want := c.firstName
		p, err := personStore.PersonByOID(c.oid)
		if err != nil {
			t.Errorf("PersonByOID(%q) err = %s", arg, err)
			continue
		}
		got := p.FirstName
		if got != want {
			t.Errorf("PersonByIOD(%q) = %q, want %q", arg, got, want)
		}
	}
}

func testPeople(t *testing.T) {
	xp, err := personStore.People()
	if err != nil {
		t.Fatalf("People() err = %s", err)
	}
	// check count
	want := 5
	got := len(xp)
	if got != want {
		t.Errorf("People() count = %d, want %d", got, want)
	}
}
