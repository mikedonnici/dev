# httpservice app

## Overview

This is a demo http service inspired by this
[Matt Ryer article](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831).

It builds on the datastore idea from the MappCPD project - a datastore
being a simple struct that holds connections to one or more data sources.

In this case, however, the functions that fetch and format data are
methods on the `datastore`. This makes testing easier and means that the
`datastore` is responsible for all aspects of data storage and retrieval
for the `httpservice`.

Integration tests are run against *real* databases which are set up from
the `testdata` dir.

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

These can be set in three ways, in order of precendence:

Firstly, by specifying a config file with the `-c` flag, eg:

```
# go run main -c "./env_example.txt"
```

Secondly, in the absence of a specified config file it will look for
the default `.env` file.

Finally, if the deployment environment allows for env vars to be set via
a control panel or similar (eg Heroku)then no configuration file needs
to be specified.

## Testing

Most of the integration tests are run against real databases with a
small set of data.

To run all tests from root dir:

```
# go test ./...
```

Test files can be run individually as each `*_test.go` file sets up its
own test database and then runs a group  of test in parallel (see
`datastore/person_test.go`).

To run an individual set of tests:

```
# go test -v person_test.go
```

Ref: https://blog.golang.org/subtests

