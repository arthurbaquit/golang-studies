package main

import "fmt"

type employee struct {
	annualIncome   int
	name, position string
}

func NewEmployeeFactory(annualIncome int, position string) func(name string) *employee {
	return func(name string) *employee {
		return &employee{annualIncome, name, position}
	}
}

func main() {
	eFac := NewEmployeeFactory(100000, "Developer")
	e := eFac("Arthur")
	fmt.Println(e)
	mFac := NewEmployeeFactory(120000, "Manager")
	m := mFac("Baquit")
	fmt.Println(m)
}
