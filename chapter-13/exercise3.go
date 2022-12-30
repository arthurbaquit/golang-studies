package main

func main(){
	defer func(){
		println("I'm the first defer")
	}()
	defer func(){
		println("I'm the second defer")
	}()
	defer func(){
		println("I'm the third defer")
	}()
	println("I'm the main function")
}