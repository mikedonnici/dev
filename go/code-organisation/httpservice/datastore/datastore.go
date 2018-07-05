// Package datastore provides a common structure for accessing data across multiple database connections
package datastore

import (
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mongo"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mysql"
)

// Datastore contains connections to the data sources required for the service
type Datastore struct {
	MySQL *mysql.Connection
	Mongo *mongo.Connection
}

// New returns a pointer to a Datastore
func New() *Datastore {
	return &Datastore{}
}
