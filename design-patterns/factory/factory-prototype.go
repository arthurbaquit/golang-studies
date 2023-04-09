package main

type employee struct {
	annualIncome   int
	Name, position string
}

const (
	Developer = iota
	Manager
)

func NewEmployeeFactoryFunc(position int) func(name string) *employee {
	switch position {
	case Developer:
		return func(name string) *employee {
			return &employee{100000, name, "Developer"}
		}
	case Manager:
		return func(name string) *employee {
			return &employee{120000, name, "Manager"}
		}
	default:
		panic("unsupported position")
	}
}

func NewEmployeeFactoryStruct(position int) *employee {
	switch position {
	case Developer:
		return &employee{100000, "", "Developer"}
	case Manager:
		return &employee{120000, "", "Manager"}
	default:
		panic("unsupported position")
	}
}

func main() {
	facFunc := NewEmployeeFactoryFunc(Developer)
	e := facFunc("Arthur")
	println(e.Name, e.position, e.annualIncome)
	facStruct := NewEmployeeFactoryStruct(Manager)
	facStruct.Name = "Baquit"
	println(facStruct.Name, facStruct.position, facStruct.annualIncome)
}
