package main

import (
	"sync"
	"time"
)

var (
	totalMoney = 0
	lock       = sync.Mutex{}
	condition  = sync.NewCond(&lock)
)

func stingy() {
	for i := 1; i < 1000; i++ {
		lock.Lock()
		totalMoney += 10
		println("Total money: ", totalMoney)

		condition.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy done")
}

func spendy() {
	for i := 1; i < 1000; i++ {
		lock.Lock()
		for totalMoney-15 < 0 {
			condition.Wait()
		}
		totalMoney -= 15
		println("Total money: ", totalMoney)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy done")

}
func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	println("Final money: ", totalMoney, "Expected: 0")
}
