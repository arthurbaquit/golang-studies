package solid

// Open-Closed Principle
//	Software entities (classes, modules, functions, etc.) should be open for
//	extension, but closed for modification.
//
//	What does this mean? It means that we should be able to extend the
//	behavior of a class without modifying it. This is a very important
//	principle, because it allows us to make changes to our code without
//	causing problems in other parts of the code.

// Let's say that we have a class Product that has a name and a color. We
// want to be able to print the product to the console, but we want to
// print the name in uppercase and the color in lowercase. We could do

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type Product struct {
	name  string
	color Color
	size  Size
}
type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, p := range products {
		if p.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// However, this is not a good solution, because if we want to add a new
// filter, we would have to modify the Filter class. This is not good,
// because we would have to modify the class every time we want to add a
// new filter. We should be able to add new filters without modifying the
// Filter class. We can do this by creating an interface that defines the
// behavior of a filter, and then we can create a new class that implements
// this interface. This is the Open-Closed Principle in action.

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct{}

type AndSpecification struct {
	first, second Specification
}

func (a *AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}
	return result
}

// In this way, the Filter class is closed for modification, but it is open
// for extension. We can add new filters without modifying the Filter class,
// by adding new classes that implement the Specification interface.
func mainOCP() {
	apple := Product{"Apple", Red, Small}
	tree := Product{"Tree", Green, Large}
	house := Product{"House", Blue, Large}

	products := []Product{apple, tree, house}
	f := Filter{}
	for _, p := range f.FilterByColor(products, Green) {
		println(p.name, "is green")
	}

	green := ColorSpecification{Green}
	large := SizeSpecification{Large}
	bf := BetterFilter{}
	for _, p := range bf.Filter(products, &green) {
		println(p.name, "is green")
	}
	for _, p := range bf.Filter(products, &large) {
		println(p.name, "is large")
	}
	and := AndSpecification{&green, &large}
	for _, p := range bf.Filter(products, &and) {
		println(p.name, "is green and large")
	}

}
