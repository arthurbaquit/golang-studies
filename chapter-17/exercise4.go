package main

import (
	"fmt"
	"sort"
)

func main(){
	xi := []int{5, 8, 2, 43, 17, 982, 124, 532}
	xs := []string{"great", "random", "rainbow", "test", "delights", "fancy", "experience","height"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}