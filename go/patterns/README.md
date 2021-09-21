# Patterns and Examples

Some useful patterns and examples adapted from Bodner and Titmus books.

## The Done Channel Pattern

- Returns the fastest response from a group of goroutines
- All remaining goroutines will close once the value has been read

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const slowestResponseMilliseconds = 3000
const numGoroutines = 10

type result struct {
	routineNum int
	responseMS int
}

func main() {

	fmt.Println("Go routines:")

	data := make(chan result)
	done := make(chan struct{})

	for i := 0; i < numGoroutines; i++ {
		i := i // need to shadow this var
		go func() {
			select {
			case data <- d(i): // 2. A result sent, this goroutine finishes
			case <-done: // 4. Read succeeds when done channel closed in main, this goroutine also finishes
			}
		}()
	}
	res := <-data // 1. Unbuffered channel blocks...
	fmt.Printf("Quickest response was on channel %d (%dms)\n", res.routineNum, res.responseMS)
	close(done) // 3. Close done channel
}

func d(n int) result {
	rand.Seed(time.Now().UnixNano())
	tms := rand.Intn(slowestResponseMilliseconds)
	fmt.Printf("d(%d) = %dms\n", n, tms)
	time.Sleep(time.Duration(tms) * time.Millisecond)
	return result{n, tms}
}
```

## Backpressure

- Use buffered channel to limit work



 
  
 
 
 
 

