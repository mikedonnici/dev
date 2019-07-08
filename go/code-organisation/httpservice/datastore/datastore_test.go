package datastore_test

import (
	"testing"

	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/datastoretest"
)

var store *datastore.Datastore

func TestDatastore(t *testing.T) {

	var teardown func()
	store, teardown = datastoretest.Setup(t)
	defer teardown()

	t.Run("datastore", func(t *testing.T) {
		t.Run("pingMySQL", testPingMySQL)
	})
}

func testPingMySQL(t *testing.T) {
	err := store.MySQL.Session.Ping()
	if err != nil {
		t.Fatalf("Ping() err = %s", err)
	}
}
