package mysql

import (
	"log"
	"testing"

	"github.com/matryer/is"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/testdata"
)

var data = testdata.New()

// TestMain sets up test databases against which subsequent datastore tests can be run. These are torn
// down once all of the tests have completed.
func TestMain(m *testing.M) {

	err := data.SetupMySQL()
	if err != nil {
		log.Fatalln(err)
	}
	defer data.TearDownMySQL()

	m.Run()
}

func TestNewConnection(t *testing.T) {
	is := is.New(t)
	_, err := NewConnection(testdata.MySQLDSN, "test", "test mysql db")
	is.NoErr(err)
}
