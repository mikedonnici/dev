// Package personsrvc provides a demo http service
package personsrvc

import (
	"errors"
)

const (
	ErrMongoDBNameEmpty = "MongoDBName not set"
	ErrMongoDSNEmpty    = "MongoDSN not set"
	ErrMySQLDBNameEmpty = "MySQLDBName not set"
	ErrMySQLDSNEmpty    = "MySQLDSN not set"
	ErrPortEmpty        = "Port not set"
)

// Config contains all of the required values to start the service
type Config struct {
	MySQLDSN    string
	MySQLDBName string
	MongoDSN    string
	MongoDBName string
	Port        string
}

// Connect to a database
func Connect(cfg Config) (Datastore, error) {

	var err error
	var ds Datastore

	err = cfg.check()
	if err != nil {
		return ds, err
	}

	ds.MySQL, err = NewMySQLConnection(cfg.MySQLDSN, cfg.MySQLDBName)
	if err != nil {
		return ds, err
	}

	ds.Mongo, err = NewMongoConnection(cfg.MongoDSN, cfg.MongoDBName)
	if err != nil {
		return ds, err
	}

	return ds, nil
}

// HTTPServer starts the http server with the attached store
func HTTPServer(port string, store Datastore) error {
	return NewServer(port, store).Start()
}

// check verifies the presence of Config field values
func (cfg *Config) check() error {

	if cfg.MongoDBName == "" {
		return errors.New(ErrMongoDBNameEmpty)
	}
	if cfg.MongoDSN == "" {
		return errors.New(ErrMongoDSNEmpty)
	}
	if cfg.MySQLDBName == "" {
		return errors.New(ErrMySQLDBNameEmpty)
	}
	if cfg.MySQLDSN == "" {
		return errors.New(ErrMySQLDSNEmpty)
	}
	if cfg.Port == "" {
		return errors.New(ErrPortEmpty)
	}

	return nil
}
