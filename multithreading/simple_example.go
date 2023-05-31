package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	now := time.Now()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100000; j++ {
		}
		for k := 0; k < 100000; k++ {
		}
		for l := 0; l < 100000; l++ {
		}
	}
	fmt.Println("Time taken: ", time.Since(now))
	now = time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(3)
		go func() {
			for j := 0; j < 100000; j++ {
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < 100000; j++ {
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < 100000; j++ {
			}
			wg.Done()
		}()
		wg.Wait()
	}
	fmt.Println("Time taken: ", time.Since(now))
}
