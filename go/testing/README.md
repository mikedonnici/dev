# Testing

*Notes from Jon Calhoun's testing with Go*

Tests do the following:

* Help find and fix mistakes, bugs, edge cases, side effects etc.
* Document expected behaviour
* Encourage better code
* Speed up development

Good tests have a **clear purpose**, such as:

* New or refactored code works as expected
* Verifies that a resolved bug does *stays* resolved 
* Documents use cases, especially for more complex APIs 


The go test tool compiles the code into an executable which is stored in a tmp folder, 
and cleaned up after the tests have completed.

### Func naming conventions:

```go
func TestType_method(t *testing.T){}
```

### Var naming conventions:

```go
arg := "argument to func being tested"
want := "expected result"
got := "actual result"
```

### Signalling failure

`t.Log() / t.Logf()` will log output to screen if:

1. A test fails
1. The test tool is run with `-v`

`t.Fail()` marks a test as failed but continues, `t.FailNow()` marks test as failed and stops.

Generally don't use in favour of the following:

* `t.Error() / t.Errorf()`- equivalent to `t.Log() / t.Logf()` + `t.Fail()`
* `t.Fatal() / t.Fatalf()` - equivalent to `t.Log() / t.Logf()` + `t.FailNow()`

## Error vs Fatal

A test run should be abandoned, ie `t.Fatal()` when there is no point continuing. Otherwise, using `t.Error()` 
is ok assuming that the subsequent tests makes sense.

For example, if a test needs to unmarshall a response body, and the response is not received, then there is 
no point continuing.

If the test involves validating parts of the response, such as headers or field values in json, then an 
 `t.Error()` might be better as other tests can continue.
 
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
The amount of detail about the args that were passed to the function will depend 
on the necessity or complexity of the function call.

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

Examples are run as tests and can be included in *_test.go files. They can also 
be viewed in teh generated documentation. 

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

For more complex examples, where it is desirable to show the entire source code including import etc, 
create a *separate* test file such as: `example_foobar_test.go` or `foobar_example_test.go`.

The generated documentation example will be complete if:

1. The example test file ends in `_test.g`
1. The example test file contains only ONE function
1. It should declare one other var, constant or function - even if it is unused.

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

Parallel subtests can be wrapped in a `t.Run()` to ensure the entire *group* of subtests complete 
before any teardown functions are called. For example:

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
Here we can see that by the time the parallel tests start to run from the queue, the value of the 
current test case is `{arg: 3, want: 9}`. So this case is run three times. The test cases all reference 
the *current* value of `tc` because of the closure.

To remedy this the easiest way is to *shadow* the variable, or create a copy, as shown below:

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
		tc := tc // create shadow (local copy) of tc for closure
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
```

See the example [parallelgotcha](./testwithgo/parallelgotcha/)

## Race Conditions

See: https://blog.golang.org/race-detector

Example of a data race: https://play.golang.org/p/02IenV-ANmT

Adding `-race` flag to most go commands will invoke the race checker.

```go
$ go test -race
$ go run -race main.go
```

This may not always catch a race condition. For example, when databases are involved.

May have to create special code to test specific race conditions.


## Testing and Value Comparisons

ref: https://golang.org/ref/spec#Comparison_operators

Simple comparisons between the same types can just use standard operators: `==`, `!=`

Need to be mindful when comparing pointer values and references.

For more complex data structures is handy to use [`reflect.DeepEqual()`](https://golang.org/pkg/reflect/#DeepEqual)


## Golden Files

Where creating a `want` value is not practical due to file size or complexity, *golden files* might be a good option.

For example, testing images or a large csv output.


## Helper Functions

When comparing values that require multiple steps to access, helper function can be useful.

For example, checking for certain values in html nodes or response headers and so on.

Some good examples: https://golang.org/src/net/http/httptest/recorder_test.go

Helper functions are also particularly useful for repeated set up and tear down processes that are required for multiple tests.

If helper functions become big enough they can introduce their own bugs so may be worthwhile creating test helpers in their own packages.

[`testing.T.Helper()`](https://golang.org/pkg/testing/#T.Helper) signals to the testing package that a function is a helper function. 












































  
   



  









 


