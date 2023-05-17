package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func goroutinePrint(i int) {
	fmt.Println("Hello from goroutine", i)
	wg.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go goroutinePrint(i)
	}
	wg.Wait()
}
