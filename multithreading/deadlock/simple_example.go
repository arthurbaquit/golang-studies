package main

import (
	"sync"
	"time"
)

var (
	lockA = sync.Mutex{}
	lockB = sync.Mutex{}
)

func First() {
	for {
		println("First asking lockA")
		lockA.Lock()
		println("First asking lockB")
		lockB.Lock()
		println("First got both locks")
		lockA.Unlock()
		lockB.Unlock()
		println("First done")
	}
}

func Second() {
	for {
		println("Second asking lockB")
		lockB.Lock()
		println("Second asking lockA")
		lockA.Lock()
		println("Second got both locks")
		lockA.Unlock()
		lockB.Unlock()
		println("Second done")
	}
}

func main() {
	go First()
	go Second()
	// Sleep for 10 seconds so that the goroutines have time to run
	time.Sleep(10 * time.Second)
	println("Done")
}
