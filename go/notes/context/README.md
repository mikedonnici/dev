# context package

The `context` package is used to coordinate the cancellation of service requests after a timeout or some other reason.

A `context` is created and passed to a service request, which in turn can pass it along to other requests.

The original context can be passed down, or a new context can be derived from it and passed along.

When a context is cancelled the function where it was derived, and all subsequent functions are cancelled, in a
coordinated way.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// A background context never gets cancelled
	ctx := context.Background()
	thingOne(ctx)
	fmt.Println("...back in main()")
}

func thingOne(ctx context.Context) {
	fmt.Println("thingOne() creates a derived context with a timeout, and passes it to thingTwo()...")
	dctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := thingTwo(dctx); err != nil {
		fmt.Println(err)
	}
	fmt.Println("...back in thingOne()")
}

func thingTwo(ctx context.Context) error {
	fmt.Println("Arrived in thingTwo(), hang here until the context times out...")
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
```

Run in [playground](https://play.golang.org/p/gixt191fcN-)