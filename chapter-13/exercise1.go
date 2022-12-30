package main

func sum (x, y int) int {
	return x + y
}

func isTheSumBiggerThanTen (x, y int) (int, string) {
	sum := sum(x, y)
	if sum > 10 {
		return sum, "yes!"
	}
	return sum, "no!"
}
func main () {
	println(sum(2, 3))
	println(isTheSumBiggerThanTen(2, 3))
	println(isTheSumBiggerThanTen(10, 3))

}
