package main

// Another approach is, instead of using a creational way to create the builder,
// we can use a function to create the builder. This approach relies on appending
// each action to a slice of functions. The slice of functions is then executed
// in order to build the object. This approach is more verbose, but it is easier
// to understand and extend. The following code is an example of this approach:

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)

type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func main() {
	pb := PersonBuilder{}
	p := pb.Called("Arthur").WorksAsA("Developer").Build()
	fmt.Println(*p)
}
