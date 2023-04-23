package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	x   int32
	wg  sync.WaitGroup
	xAt atomic.Int32
)

func addOneToVar() {
	defer wg.Done()
	runtime.Gosched()
	atomic.AddInt32(&x, 1)

}

func addOneToVarUsingAtomicVariable() {
	defer wg.Done()
	runtime.Gosched()
	xAt.Add(1)
}
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go addOneToVar()
		go addOneToVarUsingAtomicVariable()
	}
	wg.Wait()
	fmt.Println(x, xAt.Load())
}
