package main

func main(){
	println(returnSum()(1,2))

}

func returnSum() func(int, int) int {
	return func (x,y int) int{
		return x + y
	}
}