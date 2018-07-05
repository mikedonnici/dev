package main

import (
	"log"
	"os"

	"github.com/mikedonnici/dev/go/code-organisation/httpservice/server"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
	"github.com/34South/envr"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mysql"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mongo"
)

func init() {
	envr.New("appEnv", []string{
		"MYSQL_DSN",
		"MYSQL_DBNAME",
		"MYSQL_DESC",
		"MONGO_DSN",
		"MONGO_DBNAME",
		"MONGO_DESC",
	}).Auto()
}

func main() {

	var err error

	d := datastore.New()
	d.MySQL, err = mysql.NewConnection(
		os.Getenv("MYSQL_DSN"),
		os.Getenv("MYSQL_DBNAME"),
		os.Getenv("MYSQL_DESC"),
	)
	if err != nil {
		log.Fatalf("Datastore could not connect to MySQL")
	}

	d.Mongo, err = mongo.NewConnection(
		os.Getenv("MONGO_DSN"),
		os.Getenv("MONGO_DBNAME"),
		os.Getenv("MONGO_DESC"),
	)
	if err != nil {
		log.Fatalf("Datastore could not connect to MongoDB")
	}

	port := "8080" // get from env
	srv := server.NewServer(port, d)
	log.Println("server listening on port " + port)
	log.Fatal(srv.Start())
}