package main

import (
	"fmt"
	"time"
)

// Demonstrates a simple race condition - a situation where 
// multiple routines have access to a common resource. As the order
// and/or time to complete is unknown for each routine, so to is the 
// output - indeterminate routine schedule means the value of x is 
// also indeterminate.
func main() {

	// x initialised to 0
	var x int

	// first goroutine sets x = 1 after 1ms
	go func() {
		time.Sleep(1 * time.Millisecond)
		x = 1
	}()

	// second goroutine sets x = 2 after 1ms
	go func() {
		time.Sleep(1 * time.Millisecond)
		x = 2
	}()

	// print x after 1 ms - indeterminent value of x
	// as both goroutine will access as in an unknown 
	// order - a race!
	time.Sleep(1 * time.Millisecond)
	fmt.Println(x)
}