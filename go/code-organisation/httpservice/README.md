# httpservice app

This directory contains a demo http application inspired by this
[Matt Ryer article](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831).

It builds on the datastore idea from the MappCPD project - a datastore
being a struct that holds connections to one or more data sources.

In this case, however, the functions to fetch records have been hung off
the datastore to make testing easier. Thus is is the responsibility of the
`datastore` to manage all aspects of data storage and retrieval for the `httpservice`.

Testing is done against *real* databases set up from the `testdata` dir.











