package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type PressureValve struct {
	tokens      chan struct{} // read by goroutines and returned when finished
	maxTokens   int           // Number of tokens is number of goroutines that can run simultaneously
	delayMS     int           // Wait this long before retrying to get a token
	maxAttempts int           // Give up after this many attempts to fetch a token
}

func NewPressureValve(maxGoroutines, delayMS, maxAttempts int) PressureValve {

	// Fill up the tokens
	ch := make(chan struct{}, maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		ch <- struct{}{}
	}

	return PressureValve{
		tokens:      ch,
		delayMS:     delayMS,
		maxAttempts: maxAttempts,
	}
}

// Call attempts to get a work token and if successful, will invoke the function passed in.
func (v PressureValve) Call(f func()) error {

	for i := 0; i < v.maxAttempts; i++ {
		select {
		case <-v.tokens:
			fmt.Printf("Got a token!\n")
			f()
			v.tokens <- struct{}{} // return the token
			return nil
		default:
			fmt.Printf("No token available, will try again in %dms\n", v.delayMS)
		}
		time.Sleep(time.Duration(v.delayMS) * time.Millisecond)
	}
	fmt.Println("failed")
	return fmt.Errorf("failed to get a token after %d attempts", v.maxAttempts)
}

func main() {

	f := func() {
		delayMS := rand.Intn(2000) + 1000 // can take 1 - 3 seconds
		time.Sleep(time.Duration(delayMS) * time.Millisecond)
		log.Printf("f() ran - took %d\n", delayMS)
	}
	pv := NewPressureValve(10, 1000, 3)

	for i := 0; i < 100; i++ {
		fmt.Printf("Creating goroutine index %d\n", i)
		i := i
		go func() {
			err := pv.Call(f)
			if err != nil {
				fmt.Printf("goroutine index %d, err = %s\n", i, err)
			}
		}()
	}

	// Hang here for a bit so we can see what happened.
	time.Sleep(20 * time.Second)
}
