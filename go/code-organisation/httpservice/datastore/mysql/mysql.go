package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// Connection holds values for a connection to a MySQL server
type Connection struct {
	DSN     string // Data source name, format: user:pass@tcp(hostname:3306)/ - exclude database name
	DBName  string // Database name
	Desc    string // Data source description to differentiate multiple sources
	Session *sql.DB
}

// NewConnection returns a pointer to a Connection
func NewConnection(dsn, dbname, desc string) (*Connection, error) {
	m := &Connection{
		DSN:    dsn,
		DBName: dbname,
		Desc:   desc,
	}
	err := m.checkFields()
	if err != nil {
		return m, err
	}

	m.Session, err = sql.Open("mysql", m.DSN + m.DBName)
	return m, err
}

// Close terminates the Session - don't really need?
//func (m *Connection) Close() {
//	m.Session.Close()
//}

func (m *Connection) checkFields() error {
	if m.DSN == "" {
		return errors.New("MySQL DSN (data source name / connection string) not set")
	}
	if m.DBName == "" {
		return errors.New("MySQL db name not set")
	}
	if m.Desc == "" {
		return errors.New("MySQL desc not set")
	}
	return nil
}