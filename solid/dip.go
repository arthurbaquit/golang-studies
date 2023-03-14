package solid

import "fmt"

// Dependency Inversion Principle
//
// The Dependency Inversion Principle (DIP) states that high-level modules should
// not depend on low-level modules. Both should depend on abstractions. Abstractions
// should not depend on details. Details should depend on abstractions.

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level model (storage mechanism)
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level model (business logic)
type Research struct {
	relationships Relationships
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Printf("Found %s which is John's child\n", rel.to.name)
		}
	}
}

// Notice that this model breaks DIP, because Research depends on a low-level
// module (Relationships). We can fix this by creating an interface that
// represents the behavior of the low-level module, and then we can pass this
// interface to the high-level module. This way, the high-level module depends
// on an abstraction, and not a low-level module.

// This is the Dependency Inversion Principle in action, by creating an abstractions
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}
type Research2 struct {
	// relationships Relationships
	browser RelationshipBrowser
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Research2) Investigate2() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Printf("John has a child called %s\n", p.name)
	}
}

func mainDIP() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{relationships}
	research.Investigate()
	research2 := Research2{&relationships}
	research2.Investigate2()
}
