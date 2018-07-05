package main

import (
	"log"
	"os"
	"flag"
	"fmt"

	"github.com/mikedonnici/dev/go/code-organisation/httpservice/server"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore"
	"github.com/34South/envr"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mysql"
	"github.com/mikedonnici/dev/go/code-organisation/httpservice/datastore/mongo"
)

const defaultPort = "8080"

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

	portFlag := flag.String("p", "", "Specify port number (optional)")
	flag.Parse()
	fmt.Println(*portFlag)

	// Setup datastore
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

	// Set port number
	port := setPort(*portFlag)
	srv := server.NewServer(port, d)

	// Start server
	log.Println("server listening on port " + port)
	log.Fatal(srv.Start())
}

// setPort sets the port number for the server, with the env var taking the highest precedence.
func setPort(port string) string {
	if os.Getenv("PORT") != "" {
		return os.Getenv("PORT")
	}
	if port != "" {
		return port
	}
	return defaultPort
}