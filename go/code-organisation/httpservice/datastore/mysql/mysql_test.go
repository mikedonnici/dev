package mysql

import (
	"log"
	"testing"

	"github.com/mikedonnici/dev/go/code-organisation/httpservice/testdata"
)

var store = testdata.New()

// TestMain sets up test databases against which subsequent tests can be run.
// These are torn down when the tests have completed.
func TestMain(m *testing.M) {

	err := store.SetupMySQL()
	if err != nil {
		log.Fatalln(err)
	}
	defer store.TearDownMySQL()

	m.Run()
}

func TestNewConnection(t *testing.T) {
	_, err := NewConnection(testdata.MySQLDSN, store.DBName)
	if err != nil {
		t.Fatalf("NewConnection err = %s", err)
	}
}
