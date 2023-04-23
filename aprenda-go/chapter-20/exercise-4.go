package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	x     int
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func addOneToVar() {
	defer wg.Done()
	runtime.Gosched()
	mutex.Lock()
	x++
	mutex.Unlock()
}
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addOneToVar()
	}
	wg.Wait()
	fmt.Println(x)
}
