package main

import "github.com/mikedonnici/dev/practice/kata/go/httpservice/webapp"

func main() {
	srv := webapp.New("8080")
	srv.Start()
}
