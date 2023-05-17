package main

import (
	"fmt"
)

func main() {
	// cs := make(chan<- int) this produce error because the channel is receiving only
	cs := make(chan int) // this is a bidirectional channel
	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
