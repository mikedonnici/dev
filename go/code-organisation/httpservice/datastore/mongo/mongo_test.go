package mongo

import (
	"log"
	"testing"

	"github.com/matryer/is"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/testdata"
)

var data = testdata.New()

func TestMain(m *testing.M) {

	err := data.SetupMongoDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer data.TearDownMongoDB()

	m.Run()
}

func TestNewConnection(t *testing.T) {
	is := is.New(t)
	_, err := NewConnection(testdata.MongoDSN, "test", "test mongo db")
	is.NoErr(err)
}

