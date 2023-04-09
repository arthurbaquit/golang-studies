package main

import "fmt"

type employee struct {
	annualIncome   int
	name, position string
}

type employeeFactoryImpl struct {
	annualIncome int
	position     string
}

func (e *employeeFactoryImpl) createEmployee(name string) *employee {
	return &employee{e.annualIncome, name, e.position}
}

func main() {
	eFac := &employeeFactoryImpl{annualIncome: 100000, position: "Developer"}
	e := eFac.createEmployee("Arthur")
	fmt.Println(e)
	mFac := &employeeFactoryImpl{annualIncome: 120000, position: "Manager"}
	m := mFac.createEmployee("Baquit")
	fmt.Println(m)
}
