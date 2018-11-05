package main

import (
	"fmt"
	"sync"
)

var balance = 100

func main() {
	fmt.Println("This has a race condition")
	var wg sync.WaitGroup
	wg.Add(2)
	go Subtract(10, &wg)
	go Subtract(20, &wg)
	go Subtract(45, &wg)
	wg.Wait()
	fmt.Println("Balance =", balance)
}

func Subtract(amount int, wg *sync.WaitGroup) {
	balance = balance - amount
	wg.Done()
}
