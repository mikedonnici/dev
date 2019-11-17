package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// MySQLConnection holds values for a connection to a MySQL server
type MySQLConnection struct {
	DSN     string // Data source name, format: user:pass@tcp(hostname:3306)/ - exclude database name
	DBName  string // Database name
	Session *sql.DB
}

// NewMySQLConnection returns a pointer to a Connection
func NewMySQLConnection(dsn, dbName string) (*MySQLConnection, error) {
	m := &MySQLConnection{
		DSN:    dsn,
		DBName: dbName,
	}
	err := m.checkFields()
	if err != nil {
		return m, err
	}

	m.Session, err = sql.Open("mysql", m.DSN + m.DBName)
	return m, err
}

func (m *MySQLConnection) checkFields() error {
	if m.DSN == "" {
		return errors.New("MySQL DSN (store source name / connection string) not set")
	}
	if m.DBName == "" {
		return errors.New("MySQL db name not set")
	}
	return nil
}
