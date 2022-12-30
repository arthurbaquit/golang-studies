package main

func sum (x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func sumWithSlice (x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func main () {
	println(sum(2, 3))
	println(sum([]int{2, 3, 4, 5, 6, 7, 8, 9, 10}...))
	println(sumWithSlice([]int{2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
