package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	ch := make(chan string)
	quit := make(chan int)
	var rec string
	var abort bool

	for {
		if abort {
			break
		}

		go userInput(ch, quit)

		select {
		case rec = <-ch:
			fmt.Println("Received on ch: ", rec)
		case <-quit:
			abort = true
		}
	}

	fmt.Println("Received a signal on the quit channel")
}

func userInput(ch chan string, quit chan int) {
	fmt.Print("Enter string, 'q' to quit: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	if s == "q" {
		quit <- 1
		return
	}

	ch <- s
}
