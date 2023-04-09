package main

import "fmt"

type person struct {
	age   int
	name  string
	isOld bool
}

type Person interface {
	SayHello()
}

func (p *person) SayHello() {
	fmt.Println("Hello, my name is", p.name)
}

type oldPerson struct {
	age   int
	name  string
	isOld bool
}

func (p *oldPerson) SayHello() {
	fmt.Println("Hello, my name is", p.name, "and back in the days, the life was better.")
}

// This is called factory function and is used to create the struct. Notice that
// with this approach, the type is not exposed to the client, only the factory func.
// This approach is useful when the struct has logic to decide which struct to create or some default values.
func CreatePerson(name string, age int) Person {
	if age > 60 {
		return &oldPerson{age, name, true}
	}
	return &person{age, name, false}
}

func main() {
	p := CreatePerson("Arthur", 26)
	p.SayHello()
	p = CreatePerson("Baquit", 70)
	p.SayHello()
}
