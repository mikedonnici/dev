// Package datastoretest provides helper functions for datastore tests
package datastoretest

import (
	"log"
	"testing"

	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mongo"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mysql"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/testdata"
)

// Setup creates and returns a new datastore and teardown function
func Setup(t *testing.T) (*datastore.Datastore, func()) {
	t.Helper()

	// The job of testStore is only to create the new test databases
	testStore := testdata.New()
	err := testStore.SetupMySQL()
	if err != nil {
		log.Fatalf("TestStore.SetupMySQL() err = %s", err)
	}
	err = testStore.SetupMongoDB()
	if err != nil {
		log.Fatalf("TestStore.SetupMongoDB() err = %s", err)
	}

	// Create a datastore and connect it to the newly created test databases
	ds := datastore.New()
	ds.MySQL, err = mysql.NewConnection(testdata.MySQLDSN, testStore.DBName)
	if err != nil {
		log.Fatalf("mysql.NewConnection(%s, %s) err = %s", testdata.MySQLDSN, testStore.DBName, err)
	}
	ds.Mongo, err = mongo.NewConnection(testdata.MongoDSN, testStore.DBName)
	if err != nil {
		log.Fatalf("mongo.NewConnection(%s, %s) err = %s", testdata.MongoDSN, testStore.DBName, err)
	}


	teardown := func() {
		err := testStore.TearDownMySQL()
		if err != nil {
			t.Errorf("TearDownMySQL() err = %s", err)
		}
		err = testStore.TearDownMongoDB()
		if err != nil {
			t.Errorf("tearDownMongoDB() err = %s", err)
		}
	}

	return ds, teardown
}
