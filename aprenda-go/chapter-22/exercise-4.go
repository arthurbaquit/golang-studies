package main

import (
	"fmt"
)

func withQbeingExit() {
	q := make(chan int)
	c := gen2(q)

	receive2(c, q)

	fmt.Println("about to exit withQbeingExit")
}

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit main")
	fmt.Println("about to start withQbeingExit")
	withQbeingExit()
}

func gen(q chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				q <- i
				continue
			}
			c <- i
		}
		close(c)
		close(q)
	}()
	return c
}

func gen2(q chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		q <- 0

		close(c)
	}()
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("from the c channel:", v)
		case v := <-q:
			fmt.Println("from the q channel:", v)
			if v == 10 {
				return
			}
		}
	}
}

func receive2(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("from the c channel:", v)
		case <-q:
			return
		}
	}
}
