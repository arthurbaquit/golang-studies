package main

import (
	"fmt"
	"time"
)

func Step1(step1, step2 chan int) {
	for val := range step1 {
		for i := 0; i < 100000; i++ {
		}
		step2 <- 2 * val
	}
	close(step2)
}

func Step2(step2, step3 chan int) {
	for val := range step2 {
		for i := 0; i < 100000; i++ {
		}
		step3 <- 2 * val
	}
	close(step3)
}

func Step3(step3 chan int, res chan []int) {
	result := []int{}
	for val := range step3 {
		for i := 0; i < 100000; i++ {
		}
		result = append(result, val)
	}
	res <- result
	close(res)
}

func main() {
	fmt.Println("sem channels")
	result1 := []int{}
	begin := time.Now()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100000; j++ {
		}
		for k := 0; k < 100000; k++ {
		}
		for l := 0; l < 100000; l++ {
		}
		result1 = append(result1, 4*i)
	}
	fmt.Println("Time taken: ", time.Since(begin))
	fmt.Println(result1)

	fmt.Println("com channels")
	begin = time.Now()
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
	fmt.Println("Time taken: ", time.Since(begin))
	fmt.Println(<-res)
}
