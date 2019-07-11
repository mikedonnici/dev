package personsrvc_test

import (
	"testing"

	"github.com/mikedonnici/dev/go/code-organisation/httpserviceflat/personsrvc"
	"github.com/mikedonnici/dev/go/code-organisation/httpserviceflat/personsrvc/testdata"
)

var store personsrvc.Datastore

func TestDatastore(t *testing.T) {

	var teardown func()
	store, teardown = testdata.Setup(t)
	defer teardown()

	t.Run("datastore", func(t *testing.T) {
		t.Run("pingMySQL", testPingMySQL)
		t.Run("pingMongo", testPingMongo)
	})
}

func testPingMySQL(t *testing.T) {
	err := store.MySQL.Session.Ping()
	if err != nil {
		t.Fatalf("MySQL.Session.Ping() err = %s", err)
	}
}

func testPingMongo(t *testing.T) {
	err := store.Mongo.Session.Ping()
	if err != nil {
		t.Fatalf("Mongo.Session.Ping() err = %s", err)
	}
}
