# Go Modules

Extract from <https://blog.golang.org/using-go-modules>

- A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root
- `go.mod` defines the **module path** which is the import path for the root directory and its _depedenency requirements_
- Each dependency requirement is written as a module path and a specific [semantic version](https://semver.org)
- `go.mod` only appears in the root of the module
- `go` command resolves imports by using the specific dependency module versions listed in `go.mod`
- Defaults to _latest_ in order of preference for _latest tagged stable_, _latest tagged pre-release) then \_latest untagged_
- An untagged release with be listed with a [psuedo version](https://golang.org/cmd/go/#hdr-Pseudo_versions) number
- `go.sum` file contains hashes of the content of specific module versions
- Each major version (except v1) must be in module path, for example:

  - `github.com/mikedonnici/gothing` - all v1 versions
  - `github.com/mikedonnici/gothing/v2` - all v2 versions
  - `github.com/mikedonnici/gothing/v3` - all v3 versions

- Minor versions _should_ be backwards compatible

Create a module:

```bash
go mod init repo.com/pkgname
```

List current module (_main module_) and dependencies:

```bash
go list -m all
```

The _main module_ is always first on the list.

List available tagged versions of a module

```bash
$ go list -m -versions rsc.io/sampler rsc.io/sampler
v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99
```

Fetch a module at a specific version:

```bash
go get rsc.io/sampler@v1.3.1
```

Check docs for a major version:

```bash
go doc rsc.io/quote/v3
```

Clean up unused dependencies:

```bash
go mod tidy
```

## Summary

- `go mod init` creates a new module, initializing the go.mod file that describes it.
- `go build`, `go test`, and other package-building commands add new dependencies to `go.mod` as needed.
- `go list -m all` prints the current module's dependencies.
- `go get` changes the required version of a dependency (or adds a new dependency).
- `go mod tidy` removes unused dependencies.

## References

- [The Principles of Versioning in Go](https://research.swtch.com/vgo-principles) - Russ Cox
