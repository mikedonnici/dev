package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/34South/envr"
	"github.com/mikedonnici/dev/go/code-organisation/httpserviceflat/personsrvc"
)

const defaultPort = "8080"

func main() {

	var err error

	// flags
	portFlag := flag.String("p", "", "Specify port number (optional)")
	cfgFlag := flag.String("c", "", "Specify cfg file (optional - will override env vars)")
	flag.Parse()

	port := setPort(*portFlag)
	setEnv(*cfgFlag)

	cfg := personsrvc.Config{
		MongoDBName: os.Getenv("MONGO_DBNAME"),
		MongoDSN:    os.Getenv("MONGO_DSN"),
		MySQLDBName: os.Getenv("MYSQL_DBNAME"),
		MySQLDSN:    os.Getenv("MYSQL_DSN"),
		Port:        port,
	}

	store, err := personsrvc.Connect(cfg)
	if err != nil {
		log.Fatalf("personsrvc.Connect() err = %s", err)
	}

	// At this point we have a connected store so can run the available endpoint
	// queries directly
	p1, _ := store.PersonByID("1")
	fmt.Println(p1)

	// Start HTTP server
	err = personsrvc.HTTPServer(cfg.Port, store)
	if err != nil {
		log.Fatalf("personsrvc.HTTPServer() err = %s", err)
	}
	log.Println("server listening on port " + cfg.Port)
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

	// required env vars
	e := envr.New("appEnv", []string{
		"MYSQL_DSN",
		"MYSQL_DBNAME",
		"MONGO_DSN",
		"MONGO_DBNAME",
	})

	// use cfg file if present
	if cfg != "" {
		fmt.Println("Setting env from", cfg)
		e.Files = []string{cfg}
	}
	e.Auto()
}
