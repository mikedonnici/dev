package personsrvc_test

import (
	"testing"

	"github.com/mikedonnici/dev/go/code-organisation/httpserviceflat/personsrvc"
	"github.com/mikedonnici/dev/go/code-organisation/httpserviceflat/personsrvc/testdata"
)

// Must be unique across all tests
var personsrvcStore personsrvc.Datastore
const personsrvcPort = "8889"

func TestPersonSRVC(t *testing.T) {

	var teardown func()
	personsrvcStore, teardown = testdata.Setup(t)
	defer teardown()

	t.Run("personsrvc", func(t *testing.T) {
		t.Run("configError", testConfigError)
		t.Run("connect", testConnect)
		t.Run("connectMySQLError", testConnectMySQLError)
		t.Run("connectMongoError", testConnectMongoError)
	})
}

// check for errors with Config
func testConfigError(t *testing.T) {

	cases := []struct {
		cfg     personsrvc.Config
		wantErr string
	}{
		{
			cfg: personsrvc.Config{
				MongoDBName: "",
				MongoDSN:    "MONGO_DSN",
				MySQLDBName: "MYSQL_DBNAME",
				MySQLDSN:    "MYSQL_DSN",
				Port:        "8888",
			},
			wantErr: personsrvc.ErrMongoDBNameEmpty,
		},
		{
			cfg: personsrvc.Config{
				MongoDBName: "MONGO_DBNAME",
				MongoDSN:    "",
				MySQLDBName: "MYSQL_DBNAME",
				MySQLDSN:    "MYSQL_DSN",
				Port:        "8888",
			},
			wantErr: personsrvc.ErrMongoDSNEmpty,
		},
		{
			cfg: personsrvc.Config{
				MongoDBName: "MONGO_DBNAME",
				MongoDSN:    "MONGO_DSN",
				MySQLDBName: "",
				MySQLDSN:    "MYSQL_DSN",
				Port:        "8888",
			},
			wantErr: personsrvc.ErrMySQLDBNameEmpty,
		},
		{
			cfg: personsrvc.Config{
				MongoDBName: "MONGO_DBNAME",
				MongoDSN:    "MONGO_DSN",
				MySQLDBName: "MYSQL_DBNAME",
				MySQLDSN:    "",
				Port:        "8888",
			},
			wantErr: personsrvc.ErrMySQLDSNEmpty,
		},
		{
			cfg: personsrvc.Config{
				MongoDBName: "MONGO_DBNAME",
				MongoDSN:    "MONGO_DSN",
				MySQLDBName: "MYSQL_DBNAME",
				MySQLDSN:    "MYSQL_DSN",
				Port:        "",
			},
			wantErr: personsrvc.ErrPortEmpty,
		},
	}

	for _, c := range cases {
		_, err := personsrvc.Connect(c.cfg)
		if err == nil {
			t.Error("personsrvc.Connect() err = nil, want not nil")
		}
		got := err.Error()
		want := c.wantErr
		if got != want {
			t.Errorf("personsrvc.Connect() err = %s, want %s", got, want)
		}
	}
}

// This tests the .Connect() method on a Datastore. This is a bit circular as
// the testdata package has already created the databases and leveraged Datastore
// to do so. However, the Connect() method is explicitly tested here.
func testConnect(t *testing.T) {

	cfg := personsrvc.Config{
		MongoDBName: personsrvcStore.Mongo.DBName,
		MongoDSN:    personsrvcStore.Mongo.DSN,
		MySQLDBName: personsrvcStore.MySQL.DBName,
		MySQLDSN:    personsrvcStore.MySQL.DSN,
		Port:        "8888", // doesn't matter here
	}

	_, err := personsrvc.Connect(cfg)
	if err != nil {
		t.Fatalf("Connect() err = %s", err)
	}
}

// Test a MySQL connection error
func testConnectMySQLError(t *testing.T) {

	cfg := personsrvc.Config{
		MongoDBName: personsrvcStore.Mongo.DBName,
		MongoDSN:    personsrvcStore.Mongo.DSN,
		MySQLDBName: "DBNameIsNotUsed",
		MySQLDSN:    "BUNG_DSN",
		Port:        "8888", // doesn't matter here
	}
	_, err := personsrvc.Connect(cfg)
	if err == nil {
		t.Fatal("Connect() err == nil, want error")
	}
}

// Test a MongoDB connection error. Note that this test  will onl
func testConnectMongoError(t *testing.T) {

	// Ping error
	cfg := personsrvc.Config{
		MongoDBName: "aaa",
		MongoDSN:    "bbb",
		MySQLDBName: personsrvcStore.MySQL.DBName,
		MySQLDSN:    personsrvcStore.MySQL.DSN,
		Port:        "8888", // doesn't matter here
	}

	_, err := personsrvc.Connect(cfg)
	if err == nil {
		t.Fatal("Connect() err == nil, want error")
	}
}