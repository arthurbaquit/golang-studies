package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	x  int
	wg sync.WaitGroup
)

func addOneToVar() {
	defer wg.Done()
	runtime.Gosched()
	x++

}
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addOneToVar()
	}
	wg.Wait()
	fmt.Println(x)
}
