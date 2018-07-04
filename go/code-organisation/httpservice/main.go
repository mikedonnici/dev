package main

import (
	"log"
	"os"

	"github.com/mikedonnici/dev/go/code-organisation/httpservice/server"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
	"github.com/34South/envr"
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

	d := datastore.New()

	d.MySQL.DSN = os.Getenv("MYSQL_DSN")
	d.MySQL.DBName = os.Getenv("MYSQL_DBNAME")
	d.MySQL.Desc = os.Getenv("MYSQL_DESC")
	err := d.MySQL.Connect()
	if err != nil {
		log.Fatalf("Datastore could not connect to MySQL")
	}

	d.Mongo.DSN = os.Getenv("MONGO_DSN")
	d.Mongo.DBName = os.Getenv("MONGO_DBNAME")
	d.Mongo.Desc = os.Getenv("MONGO_DESC")
	err = d.Mongo.Connect()
	if err != nil {
		log.Fatalf("Datastore could not connect to MongoDB")
	}

	port := "8080" // get from env
	log.Println("server listening on port " + port)
	log.Fatal(server.Start(port, d))
}
