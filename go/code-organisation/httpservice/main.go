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

//func init() {
//	envr.New("appEnv", []string{
//		"MYSQL_DSN",
//		"MYSQL_DBNAME",
//		"MYSQL_DESC",
//		"MONGO_DSN",
//		"MONGO_DBNAME",
//		"MONGO_DESC",
//	}).Auto()
//}

func main() {

	var err error

	// flags
	portFlag := flag.String("p", "", "Specify port number (optional)")
	cfgFlag := flag.String("c", "", "Specify cfg file (optional - will override env vars)")
	flag.Parse()

	port := setPort(*portFlag)
	setEnv(*cfgFlag)

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

	// Start server
	srv := server.NewServer(port, d)
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

func setEnv(cfg string) {

	// declare required env vars
	e := envr.New("appEnv", []string{
		"MYSQL_DSN",
		"MYSQL_DBNAME",
		"MYSQL_DESC",
		"MONGO_DSN",
		"MONGO_DBNAME",
		"MONGO_DESC",
	})

	// use cfg file if present
	if cfg != "" {
		fmt.Println("Setting env from", cfg)
		e.Files = []string{cfg}
	}
	e.Auto()

	fmt.Println(e.V)
}