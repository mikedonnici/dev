package main

import (
	"fmt"
	"sync"
	"time"
)

const maxEats = 3

type chopstick sync.Mutex

type philosopher struct {
	Index          int
	LeftChopstick  sync.Mutex
	RightChopstick sync.Mutex
}

var philosophers []philosopher
var chopsticks []sync.Mutex
var wg sync.WaitGroup

func init() {

	for i := 0; i < 5; i++ {
		philosophers = append(philosophers, philosopher{Index: i})
		chopsticks = append(chopsticks, sync.Mutex{})
	}
}

func main() {

	for j := 0; j < maxEats; j++ {
		wg.Add(len(philosophers))
		for i := 0; i < 5; i++ {
			go philosophers[i].eat()
		}
		wg.Wait()
	}
}

func (p *philosopher) eat() {

	left := p.Index
	right := p.Index - 1
	if right < 0 {
		right = len(chopsticks) - 1
	}
	// request and wait for permission - this will lock the chopsticks
	ch := make(chan int)
	go hostPermission(&chopsticks[left], &chopsticks[right], ch)
	<-ch

	// eat for a second
	msg := "Starting to eat %d (with chopsticks %d and %d)\n"
	fmt.Printf(msg, p.Index+1, left+1, right+1)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finishing eating %d\n", p.Index + 1)
	// release the chopsticks
	chopsticks[left].Unlock()
	chopsticks[right].Unlock()
	wg.Done()
}

// hostPermission gives a philosopher permission to eat
// by getting a lock on the required chopsticks. It communicates
// the permission over a channel.
func hostPermission(leftChopstick, rightChopstick *sync.Mutex, permission chan int) {
	// these will block until gained
	leftChopstick.Lock()
	rightChopstick.Lock()
	permission <- 1
	// <- finished
	// leftChopstick.Unlock()
	// rightChopstick.Unlock()
}
