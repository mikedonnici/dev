

## Set Up

- Set up preferred path for third-party packages:

```shell
# add to $HOME/.profile
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

```shell
source $HOME/.profile
```

##  Third-Party Tools

- Install with `go install gthub.com/entity/repo@version`
- Installs to `$GOPATH/bin` 
- To update a package just reinstall it `@latest` 

## Formatting

- `go fmt` is standard
- `goimports` - enhanced version of `go fmt` which cleans up import statements

```shell
go install golang.org/x/tools/cmd/goimports@latest
goimports -l -w . # -l show files, -w in place
```

## Linting and Vetting

- Refs:
   - [Effective Go](https://golang.org/doc/effective_go)
   - [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
    
_Don't assume any of these are 100% accurate_
  
- `golint` - Enforces style, ie syntax structure

```shell
go install golang.org/x/lint/golint@latest
golint ./...
```

- `go vet` - Picks up programmatic errors, unused vars etc

```shell
go vet ./...
```

- [`golangci-lint`](https://github.com/golangci/golangci-lint)
- Can run 10-50+ linters over a project

## Makefiles

- Automates build steps
- Makefile must have tabs 
- Example:

```Makefile
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build
```

- `.PHONY` prevents drama if a folder exists with same name as the step
- Running just `make` will set goal to the `build` step
- Can also run any step on its own, eg `make lint`
- See [hello/](./hello/) for a working example
