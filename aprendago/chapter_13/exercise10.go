package main

func main(){
	firstCount := count()
	secondCount := count()
	println(firstCount())
	println(firstCount())
	println(firstCount())
	println(secondCount())
	println(secondCount())
	println(secondCount())
}

func count() func() int {
	x:=0
	return func () int{
		x++
		return x
	}
}