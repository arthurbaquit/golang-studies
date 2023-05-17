package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}

}
