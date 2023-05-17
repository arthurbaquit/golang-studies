package main

import (
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int)
	count := 0
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 10; i++ {
					channel <- rand.Intn(100)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(channel)
	}()

	for v := range channel {
		count++
		println(v)
	}

	println("total: ", count)
}
