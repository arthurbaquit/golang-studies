package main

import (
	"fmt"
	"sync"
	"time"
)

var rwlock = sync.RWMutex{}

func oneTwoThreeB() {
	rwlock.RLock()
	for i := 1; i <= 300; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	rwlock.RUnlock()
}

func StartThreadsB() {
	for i := 1; i <= 3; i++ {
		go oneTwoThreeB()
	}
	time.Sleep(1 * time.Second)
}

func main() {
	StartThreadsB()
}
