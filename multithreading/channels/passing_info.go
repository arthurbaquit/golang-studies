package main

import "fmt"

func Step1(step1, step2 chan int) {
	for val := range step1 {
		step2 <- 2 * val
	}
	close(step2)
}

func Step2(step2, step3 chan int) {
	for val := range step2 {
		step3 <- 2 * val
	}
	close(step3)
}

func Step3(step3 chan int, res chan []int) {
	result := []int{}
	for val := range step3 {
		result = append(result, val)
	}
	res <- result
	close(res)
}

func main() {
	step1 := make(chan int)
	step2 := make(chan int)
	step3 := make(chan int)
	res := make(chan []int)
	go Step1(step1, step2)
	go Step2(step2, step3)
	go Step3(step3, res)
	for i := 0; i < 100; i++ {
		step1 <- i
	}
	close(step1)
	fmt.Println(<-res)
}
