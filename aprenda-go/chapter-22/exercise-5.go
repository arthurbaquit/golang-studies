package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	if v, ok := <-c; ok {
		fmt.Println(v, ok)
	}

	close(c)

	if v, ok := <-c; !ok {
		fmt.Println(v, ok)
	}
}
