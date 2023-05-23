package main

import (
	"time"
)

var (
	totalMoney = 100
)

func stingy() {
	for i := 1; i <= 10000; i++ {
		totalMoney += 10
	}
	println("Stingy done")
}

func spendy() {
	for i := 1; i <= 10000; i++ {
		totalMoney -= 10
	}
	println("Spendy done")

}
func main() {
	go stingy()
	go spendy()
	time.Sleep(1000 * time.Millisecond)
	println("Final money: ", totalMoney, "Expected: 100")
}
