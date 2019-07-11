# http service example

## Overview

This is a demo http service inspired by this 
[Matt Ryer article](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831).

It builds on the datastore idea from the MappCPD project - a datastore 
being a simple struct that holds connections to one or more data sources.

In this case, however, the functions that fetch and format data are 
methods on the `datastore`. This makes testing easier and means that 
the `datastore` is responsible for all aspects of data storage and 
retrieval for the `httpservice`.

Integration tests are run against _real_ databases which are set up 
from the `testdata` dir.

In this example the datastore has a MySQL and a MongoDB database attached.

## Configuration

The service requires the following env vars (example values):

```
MYSQL_DSN="root:password@tcp(localhost:3306)/"
MYSQL_DBNAME="test"
MYSQL_DESC="local dev"
MONGO_DSN="mongodb://localhost"
MONGO_DBNAME="test"
MONGO_DESC="local mongo db"
```

These can be set in three ways, in order of precedence:

Firstly, by specifying a config file with the `-c` flag, eg:

```bash
go run main -c "./env_example.txt"
```

Secondly, in the absence of a specified config file it will look for 
the default `.env` file.

Finally, if the deployment environment allows for env vars to be set 
via a control panel or similar (eg Heroku) then the env var values will 
be read from the environment and no configuration file is necessary.

**Port Number**

If a `PORT` env var is present in the deployment environment then the 
server will listen on that port. This is the case for Heroku and similar.

Otherwise, port number can be specified with an optional `-p` flag, or 
left to the default of 8080.

## Testing

Integration tests are run against real databases with a small set of data.

These database are setup (and torn down) using functions provided in the 
`testdata` package which is ignored by the compiler.  

This means that on a local dev machine both MySQL and MongoDB should be 
running. Note that the database credentials for test are hard-coded in 
the `testdata` folder.

** add notes here on Travis CI setup for MySQL and MongoDB**

To run all tests from root dir:

```
# go test ./...
```

Test files can be run individually as each `*_test.go` file sets up its 
own test database and then runs a group of test in parallel (see 
`datastore/person_test.go`).

To run an individual set of tests:

```
# go test -v person_test.go
```

### Thoughts on structure

These are a few thoughts to be sanitised later on.

- A single 'service' (in the _microservice_ sense) represents a group of related
  _tasks_ pertaining to one or more related data entities.
- Each service has it's own `datastore` because the source(s) of data may be different for each service.

Possible structure for a UserService:

```dir
userService/
├── datastore/
│   ├── mysql/
│   │   └── mysql.go
│   ├── mongo/
│   │   └── mongo.go
│   ├── datastore.go
|   ├── person.go
│   ├── address.go  
├── server/
│   ├── routes.go
│   └── server.go
```

Or a flat structure:

```dir
userService/
├── address.go  
├── datastore.go
├── mongo.go
├── mysql.go
├── person.go
├── routes.go
└── server.go
```






## References 
<https://medium.com/@povilasve/go-advanced-tips-tricks-a872503ac859>
<https://blog.golang.org/subtests>
