package main

func main(){
	println(operation(1,2, returnSum))
}

func operation(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func returnSum(x,y int) int {
		return x + y
}