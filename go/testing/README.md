# Testing in Go

_Notes sourced primarily from Jon Calhoun's: Test with Go course_

Tests do the following:

- Help find and fix mistakes, bugs, edge cases, side effects etc.
- Document expected behaviour
- Encourage better code
- Speed up development

Good tests have a **clear purpose**, such as:

- New or refactored code works as expected
- Verifies that a resolved bug does _stays_ resolved
- Documents use cases, especially for more complex APIs

The go test tool compiles the code into an executable which is stored in a tmp folder, and cleaned up after the tests have completed.

## Func naming conventions

```go
func TestType_method(t *testing.T){}
```

## Var naming conventions

```go
arg := "argument to func being tested"
want := "expected result"
got := "actual result"
```

## Signalling failure

`t.Log() / t.Logf()` will log output to screen if:

1. A test fails
2. The test tool is run with `-v`

`t.Fail()` marks a test as failed but continues, `t.FailNow()` marks test as failed and stops.

Generally don't use in favour of the following:

- `t.Error() / t.Errorf()`- equivalent to `t.Log() / t.Logf()` + `t.Fail()`
- `t.Fatal() / t.Fatalf()` - equivalent to `t.Log() / t.Logf()` + `t.FailNow()`

## Error vs Fatal

A test run should be abandoned, ie `t.Fatal()` when there is no point continuing. Otherwise, using `t.Error()` is ok assuming that the subsequent tests makes sense.

For example, if a test needs to unmarshall a response body, and the response is not received, then there is no point continuing.

If the test involves validating parts of the response, such as headers or field values in json, then an `t.Error()` might be better as other tests can continue.

See [errorfail](./testwithgo/errorfail) for an example.

## Useful failure messages

**Goal is to show what went wrong, make it easy to find and fix**

For example, when a value is not as expected:

```go
t.Errorf("SomeFunc(%v) = %v; want %v", got, want)
```

When an error is encountered you might have:

```go
t.Fatalf("SomeFunc(%v) err = %v", err)
```

The amount of detail about the args that were passed to the function will depend on the necessity or complexity of the function call.

In a lot of cases it won't be necessary:

```go
func TestHandler(t *testing.T) {
    w := httptest.NewRecorder()
    r, err := http.NewRequest(http.MethodGet, "", nil)
    if err != nil {
        t.Fatalf("http.NewRequest() err = %s", err)
    }
    Handle(w,r)
    resp := w.Result()
    is resp.StatusCode != 200 {
        t.Fatalf("Handler() status = %d; want %d", resp.StatusCode, 200)
    }    
}
```

## Examples as test cases

Examples are run as tests and can be included in *_test.go files. They can also be viewed in teh generated documentation.

```go
func ExampleHello() {
     greeting := Hello("Mike")
     fmt.Println(greeting)
     // Output: Hello, Mike
 }
```

When `go test` is run the example is run as a test and the commented line...

```go
// Output: Hello, Mike

// ...OR...

// Output: 
// Hello, Mike
```

...tells the test tool what to expect.

Specify `// Unordered output:` in cases where order of result can't be controlled, eg maps, go routines.

[Example function names](https://blog.golang.org/examples) affect where and how the example is shown in the generated docs.

## More complex examples

For more complex examples, where it is desirable to show the entire source code including import etc, create a _separate_ test file such as: `example_foobar_test.go` or `foobar_example_test.go`.

The generated documentation example will be complete if:

1. The example test file ends in `_test.g`
2. The example test file contains only ONE function
3. It should declare one other var, constant or function - even if it is unused.

## Table-driven tests

A testing pattern that is commonly used where an anonymous struct is set up to range over a bunch of tests.

These can be automatically generated in some IDEs using the [gotests](https://github.com/cweill/gotests) package.

The generated test looks like this:

```go
func TestHello(t *testing.T) {
    type args struct {
        name string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Hello(tt.args.name); got != tt.want {
                t.Errorf("Hello() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

When subtests are used a single test failure does not cause the test suite to exit.

## Parallel Tests

Tests and subtest functions that include a call to `t.Paralell()` will be run in parallel.

Parallel testing isn't really for making tests fatser, more for testing concurrency, thread-safety etc.

Parallel subtests can be wrapped in a `t.Run()` to ensure the entire _group_ of subtests complete before any teardown functions are called. For example:

```go
func TestSomething(t *testing.T) {
    defer tearDown();
    t.Run("group", func(t *testing.T){
        t.Run("sub1", func(t *testing.T){
            t.Parallel()
            // run sub1...
        })
        t.Run("sub2", func(t *testing.T){
            t.Parallel()
            // run sub2...
        })
    })
}
```

Ranging over parallel subtests can cause problems when the changing value is referenced in a closure.

For example:

```go
func TestSquare(t *testing.T) {
    testCases := []struct {
        arg  int
        want int
    }{
        {1, 1},
        {2, 5}, // should fail
        {3, 9},
    }
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
            t.Parallel() // this causes subtest func call onto a queue and subtest exits
            t.Logf("arg=%d, want=%d", tc.arg, tc.want)
            if Square(tc.arg) != tc.want {
                t.Errorf("%d^2 != %d", tc.arg, tc.want)
            }
        })
    }
}
```

The output of this is:

```bash
=== RUN   TestSquare_fail
=== RUN   TestSquare_fail/arg=1
=== PAUSE TestSquare_fail/arg=1
=== RUN   TestSquare_fail/arg=2
=== PAUSE TestSquare_fail/arg=2
=== RUN   TestSquare_fail/arg=3
=== PAUSE TestSquare_fail/arg=3
=== CONT  TestSquare_fail/arg=1
=== CONT  TestSquare_fail/arg=3
=== CONT  TestSquare_fail/arg=2
--- PASS: TestSquare_fail (0.00s)
    --- PASS: TestSquare_fail/arg=1 (0.00s)
        parallelgotcha_test.go:20: arg=3, want=9
    --- PASS: TestSquare_fail/arg=3 (0.00s)
        parallelgotcha_test.go:20: arg=3, want=9
    --- PASS: TestSquare_fail/arg=2 (0.00s)
        parallelgotcha_test.go:20: arg=3, want=9
PASS
```

Here we can see that by the time the parallel tests start to run from the queue, the value of the current test case is `{arg: 3, want: 9}`. So this case is run three times. The test cases all reference the _current_ value of `tc` because of the closure.

To remedy this the easiest way is to _shadow_ the variable, or create a copy, as shown below:

func TestSquare(t _testing.T) { testCases := []struct { arg int want int }{ {1, 1}, {2, 5}, // should fail {3, 9}, } for_ , tc := range testCases { tc := tc // create shadow (local copy) of tc for closure t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t_ testing.T) { t.Parallel() // this causes subtest func call onto a queue and subtest exits t.Logf("arg=%d, want=%d", tc.arg, tc.want) if Square(tc.arg) != tc.want { t.Errorf("%d^2 != %d", tc.arg, tc.want) } }) } }

````
This time we get the expected result:

```bash
...
=== CONT  TestSquare_fail/arg=2
--- FAIL: TestSquare_fail (0.00s)
    --- PASS: TestSquare_fail/arg=1 (0.00s)
        parallelgotcha_test.go:41: arg=1, want=1
    --- PASS: TestSquare_fail/arg=3 (0.00s)
        parallelgotcha_test.go:41: arg=3, want=9
    --- FAIL: TestSquare_fail/arg=2 (0.00s)
        parallelgotcha_test.go:41: arg=2, want=5
        parallelgotcha_test.go:43: 2^2 != 5
FAIL
````

See the example [parallelgotcha](./testwithgo/parallelgotcha/)

## Race Conditions

See: <https://blog.golang.org/race-detector>

Example of a data race: <https://play.golang.org/p/02IenV-ANmT>

Adding `-race` flag to most go commands will invoke the race checker.

```go
$ go test -race
$ go run -race main.go
```

This may not always catch a race condition. For example, when databases are involved.

May have to create special code to test specific race conditions.

## Testing and Value Comparisons

ref: <https://golang.org/ref/spec#Comparison_operators>

Simple comparisons between the same types can just use standard operators: `==`, `!=`

Need to be mindful when comparing pointer values and references.

For more complex data structures is handy to use [`reflect.DeepEqual()`](https://golang.org/pkg/reflect/#DeepEqual)

## Golden Files

Where creating a `want` value is not practical due to file size or complexity, _golden files_ might be a good option.

For example, testing images or a large csv output.

## Helper Functions

When comparing values that require multiple steps to access, helper function can be useful.

For example, checking for certain values in html nodes or response headers and so on.

Some good examples: <https://golang.org/src/net/http/httptest/recorder_test.go>

Helper functions are also particularly useful for repeated set up and tear down processes that are required for multiple tests.

[`testing.T.Helper()`](https://golang.org/pkg/testing/#T.Helper) signals to the testing package that a function is a helper function and will prevent the helper's file and line information being included in the test output.

Helper functions are particularly useful for setup and teardown, for example, setting up a test database.

A useful pattern in this case to to create a setup function that returns a closure which is used to teardown once all the test sare done.

Something like this:

```go
func TestAll(t *testing.T) {
    var teardown func()
    dataStore, teardown = setup()
    defer teardown()

    t.Run("all", func(t *testing.T) {
        t.Run("testOne", testOne)
        t.Run("testTwo", testTwo)
        t.Run("testThree", testThree)
    })
}

func setup() (datastore.Datastore, func()) {
    t.Helper()

    db := testdata.New()
    err := db.SetupDB()
    if err != nil {
        log.Fatalf("db.SetupDB() err = %s", err)
    }

    return db, func() {
        err := db.TearDownDB()
        if err != nil {
            log.Fatalf("db.TearDownDB() err = %s", err)
        }
    }
}
```

This pattern is used here: <https://github.com/mikedonnici/rtcl-api/blob/master/datastore/user_test.go>

If the number or scale of helper functions becomes significant they can introduce their own bugs so may be worthwhile creating test helpers in a seperate sub package. This is also a handy for providing access to the helper functions for consumers of the package, for example: [httptest](https://golang.org/src/net/http/httptest/recorder_test.go)

Another approach might be to create a `testing.go` file in the package directory.

**ref:** <https://members.usegolang.com/twg/lessons/lesson-34>

## Running Specific Tests

Can run a specific test function, for example:

```bash
$ go test -v -run TestLog/log/testLogsByUserID
=== RUN   TestLog
=== RUN   TestLog/log
=== RUN   TestLog/log/testLogsByUserID
--- PASS: TestLog (0.12s)
    --- PASS: TestLog/log (0.00s)
        --- PASS: TestLog/log/testLogsByUserID (0.00s)
PASS
ok      github.com/mikedonnici/rtcl-api/datastore    0.126s
```

This will run the `testLogsByUserID` sub test within the `log` group, in the `TestLog` function, as shown below:

```go
func TestLog(t *testing.T) {

    t.Run("log", func(t *testing.T) {
        t.Run("testPingDB", testPingDB)
        t.Run("testAddLog", testAddLog)
        t.Run("testUpdateLog", testUpdateLog)
        t.Run("testDeleteLog", testDeleteLog)
        t.Run("testLogByID", testLogByID)
        t.Run("testLogByIDNotFound", testLogByIDNotFound)
        t.Run("testLogsByUserID", testLogsByUserID)
    })
}
```

The flag `-run TestLog/log/testLogsByUserID` behaves like a regular expresion so tests can be matches in various ways:can be used:

```bash
go test -v -run TestLog/log/testLogsByUserID # matches one test
go test -v -run TestLog/log/testLog          # matches 3 tests
go test -v -run TestLog/log                  # matches all tests
```

### Skipping Tests

The `-short` flag can be used in combination with `t.Skip()` as required.

For example:

```bash
go test -v -short
```

```go
func TestThing(t *testing.T) {
    if testing.Short() {
        t.Skip()
    }
    t.Log("Run long tests")
}
```

Could also create custom flags to control which tests are run, for example, `-integration` to run integration tests:

(Example from <https://github.com/joncalhoun/twg/blob/master/skip/flag_test.go>)

```go
var integration = false

func init() {
    flag.BoolVar(&integration, "integration", false, "run database integration tests")
}

func TestMain(m *testing.M) {
    flag.Parse()
    if integration {
        // setup integration stuff if you need to
    }
    result := m.Run()
    if integration {
        // teardown integration stuff if you need to
    }
    os.Exit(result)
}

func TestWithFlag(t *testing.T) {
    if !integration {
        t.Skip()
    }
    t.Log("Running the integration test...")
}
```

### Specifying tests with build tags

[Build tags](https://golang.org/pkg/go/build/) can also be used to specify a set of tests to include.

For example:

Integration tests for psql and mysql are separated into files: `tag_psql_test.go`, build tag `// +build psql` and `tag_mysql_test.go`, build tag `// +build mysql` respectively. The tests for each can then be specified as follows:

```bash
go test -v -tags=psql         # postgres tests
go test -v -tags=mysql        # mysql tests
go test -v -tags="psql mysql" # both
```

**ref**: <https://members.usegolang.com/twg/lessons/lesson-42>

## Verbose flag

The `-v` flag gives a more verbose output and will display output from `t.Log()`.

To trigger actions when a _verbose_ testing is running, can do:

```go
if testing.Verbose() {
 // ...
}
```

## Code Coverage

`-cover` flags gives coverage information:

```bash
go test -cover
```

`-coverprofile` stores coverage information in a file:

```bash
go test --coverprofile=cover.out
```

This file can then be used to provide information about coverage, using `go tool`:

```bash
go tool cover -func=cover.out
```

The above will show coverage info for each function, or use the `-html` flag to generate an html version:

```bash
go tool cover -html=cover.out
```

Here's a nice little trick!

Put this in `.bashrc` / `.bash_profile`:

```bash
cover () {
    t="/tmp/go-cover.$$.tmp"
    go test -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
}
```

Run `cover` from within the package directory to show the coverage in the terminal, and also the html version.

Refs:

- <https://coderwall.com/p/rh-v5a/get-coverage-of-golang-test>
- <https://stackoverflow.com/questions/10516662/how-to-measure-code-coverage-in-golang>
- <https://blog.golang.org/cover>

## Internal vs external tests

Internal tests are written within the same package name.

External tests are written as though from an _external_ package and use the packgage name convention: `pkgname_test`.

Common naming convention to separate internal and external package test code into files named `pkgname_internal_test.go` and `pkgname_test.go` respectively.

Generally, _external_ tests should be used wherever possible, because:

- Implementation agnostic - won't break if implementation changes, less _brittle_
- Helps with versioning - if tests do break then might need to bump major version
- Provides better documentation/examples about how to use the package
- Tends to lead to better package design - more user friendly

External test can sometimes help to resolve cyclical dependencies.

Using external tests means that the unexported types, methods, functions etc are _not_ available in test functions.

If it is desirable to access an unexported value in an _external_ test the following convention is often used:

- Create a filed called `export_test.go`
- Assign this file to the _internal_ package name
- Map the required values

`export_test.go`

```go
package foo

var Bar = bar
```

`foo_test.go`

```go
package foo_test

func TestFoo(t *testing.T) {
    got := Foo()
    // ...
}
```

Because this is an `_test` file the code will only be included in tests.

Default should be to create _external_ tests, so _internal_ test are only necessary when something cannot be achieved using _external_ tests.

## Global state

Makes testing more difficult so _should be avoided wherever possible_.

## Dependency injection

Dependency injection is a just a design pattern that enables more _implementation-agnostic_ code.

```go
// without dependency injection
func Demo1() {
    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    err := foo()
    if err != nil {
        logger.Println("foo() err =", err)
    }
}

// with dependency injection
func Demo2(logger *log.Logger) {
    err := foo()
    if err != nil {
        logger.Println("foo() err =", err)
    }
}
```

The above example is still quite strict as it requires a very specific paramter be passed in to `Demo2()`.

To make it _implementation-agnostic_ could pass in a logging function:

```go
func Demo3(logFn func(...interface{})) {
    err := foo()
    if err != nil {
        logFn("foo() err =", err)
    }
}
```

Taken one step further, use an interface:

```go
type Logger interface {
    Println(...interface{})
    Printf(string, ...interface{})
}

// Call with:
//  logger := log.New(...)
//  Demo4(logger)
func Demo4(logger Logger) {
    err := foo()
    if err != nil {
        logger.Println("foo() err =", err)   // or...
        logger.Printf("foo() err = %s\n", err)
    }
}
```

Can avoid having to pass the `Logger` in each time by using a struct:

```go
type Foo struct {
    Logger interface {
        Println(...interface{})
        Printf(string, ...interface{})
    }
}

func (f Foo) Demo5() {
    err := foo()
    if err != nil {
        f.Logger.Println("foo() err =", err)   // or...
        f.Logger.Printf("foo() err = %s\n", err)
    }
}
```

## DI and useful zero values / defaults

Dependency injection requires the dependencies to be passed in (obviously). If, for example, a large code base was being changed over to DI it may be desirable for the functions that use the dependencies to be able to handle a zero value.

For example:

```go
func Foo(bar Dependency) {
    if bar == nil {
        bar = DefaultBar()
    }
    // ...
}
```

## Removing global state with DI

Global variables should be avoided wherever possible and even unexported, package-level variables can create problems.

Sometimes it is desirable to simplify code by create some type of package-level default value that is accessible to all functions.

This can be done using a type, for example:

```go
type Foo struct {
    Logger interface{
        Println(...interface{})
        Printf(string, ...interface{})
    }
}

var defaultFoo Foo

func Bar() {
    defaultFoo.Bar()
}
```

DI make writing tests easier and enables implementation -agnostic design patters such as _domain-driven design_.

DI is not free and generally required more code to implement.
