# Functional Programming

Notes from `Learning Functional Programming in Go` by Lex Sheehan.

Code: <https://github.com/l3x/learn-fp-go>

**Imperative programming** uses statement that change a program's state, 
step-by-step mechanics of how a program operates.

**Declarative programming** is express in terms of the desired results 
rather than the steps to get there.

## Pure functions


- functions are first-class citizens
- always return the same result for same inputs (**idempotence**)
- no side effects
- results are not affected by external state
- variables do not change over time 

The last three points imply **referential transparency** which means the 
 function call can be replaced with its return value without affecting 
 the program's behaviour.
 
## Memoization

Memoization is an optimisation technique of caching results of expensive 
function calls. This works because pure functions return consistent results
and have no side effects.  

This recursive _Fibonacci_ example become very inefficient as the input 
number gets larger (try 50):

<https://play.golang.org/p/9xTHgRzGxnG>

```go
func main() {
	fib := fibonacci
	fmt.Printf("%v", fib(51))
}

func fibonacci(x int) int {

	if x == 0 {
		return 0
	}

	if x <= 2 {
		return 1
	}

	return fibonacci(x-2) + fibonacci(x-1)
}
```

Memoized version: 

<https://play.golang.org/p/otqyUOK5tm8>

```go
type Memoized func(int) int
var fibMem = Memoize(fib)

func main() {
	fmt.Printf("%v", FibMemoized(5))
}

func Memoize(mf Memoized) Memoized {

	cache := make(map[int]int)

	return func(key int) int {

		val, exists := cache[key]
		if exists {
			return val
		}
		cache[key] = mf(key)

		return cache[key]
	}
}

func FibMemoized(n int) int {
	return fibMem(n)
}

func fib(x int) int {

	if x == 0 {
		return 0
	}

	if x <= 2 {
		return 1
	}

	return fib(x-2) + fib(x-1)
}
```









 
    

 




