# Testing

## Test Coverage

Here's a nice little trick for showing test coverage in go.

Put this in .bash_profile:

```bash
cover () { 
    t="/tmp/go-cover.$$.tmp"
    go test -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
}
```

Then run `cover` from within the package directory. It will show the coverage 
in a terminal and also pop open the browser and show the html coverage.

Refs: 
* https://coderwall.com/p/rh-v5a/get-coverage-of-golang-test
* https://stackoverflow.com/questions/10516662/how-to-measure-code-coverage-in-golang
* https://blog.golang.org/cover
