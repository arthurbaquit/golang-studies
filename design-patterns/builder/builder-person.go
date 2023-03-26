package main

import "fmt"

type Person struct {
	// personal information
	StreetAddress, Postcode, City string

	// job information
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

// Create the builder object
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

// Create personal information builder

type PersonAddressBuilder struct {
	PersonBuilder
}

// this method allows us to switch to the personal information builder
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

// uses fluent pattern to return the same builder object
func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

// Create job information builder

type PersonJobBuilder struct {
	PersonBuilder
}

// this method allows us to switch to the job information builder
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

// uses fluent pattern to return the same builder object
func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("Rua Fortaleza 123").
		In("Brasil").
		WithPostcode("00000-000").
		Works().
		At("Git").
		AsA("Engineer").
		Earning(123456)

	person := pb.Build()
	fmt.Println(person)
}
