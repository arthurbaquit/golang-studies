package main

import "fmt"

// Abstract Factory
type IBrandFactory interface {
	CreateShoe() IShoe
	CreateBag() IBag
}

// Abstract Product A
type IShoe interface {
	SetLogo(logo string)
	SetColor(color string)
	GetLogo() string
	GetColor() string
}

// Abstract Product B
type IBag interface {
	SetLogo(logo string)
	SetModel(model string)
	GetLogo() string
	GetModel() string
}

// Concrete Product A
type Bag struct {
	logo  string
	model string
}

func (b *Bag) SetLogo(logo string) {
	b.logo = logo
}

func (b *Bag) SetModel(model string) {
	b.model = model
}

func (b *Bag) GetLogo() string {
	return b.logo
}

func (b *Bag) GetModel() string {
	return b.model
}

// Concrete Product B
type Shoe struct {
	logo  string
	color string
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) SetColor(color string) {
	s.color = color
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) GetColor() string {
	return s.color
}

// Concrete Factory A
type LuxuryBrandFactory struct{}

type LuxuryShoe struct {
	*Shoe
}

type LuxuryBag struct {
	*Bag
}

func (l *LuxuryBrandFactory) CreateShoe() IShoe {
	return &LuxuryShoe{
		Shoe: &Shoe{
			logo:  "Luxury",
			color: "white",
		}}
}

func (l *LuxuryBrandFactory) CreateBag() IBag {
	return &LuxuryBag{
		Bag: &Bag{
			logo:  "Luxury",
			model: "Cross Body",
		}}
}

// Concrete Factory B
type SportBrandFactory struct{}

type SportShoe struct {
	*Shoe
}

type SportBag struct {
	*Bag
}

func (s *SportBrandFactory) CreateShoe() IShoe {
	return &SportShoe{
		Shoe: &Shoe{
			logo:  "Sport",
			color: "red",
		}}
}

func (s *SportBrandFactory) CreateBag() IBag {
	return &SportBag{
		Bag: &Bag{
			logo:  "Sport",
			model: "Backpack",
		}}
}

func GetBrandFactory(brand string) (IBrandFactory, error) {
	switch brand {
	case "luxury":
		return &LuxuryBrandFactory{}, nil
	case "sport":
		return &SportBrandFactory{}, nil
	default:
		return nil, fmt.Errorf("Brand %s not found", brand)
	}
}

func main() {
	factory, _ := GetBrandFactory("sport")
	shoe := factory.CreateShoe()
	bag := factory.CreateBag()
	printDetails(shoe, bag)

	factory, _ = GetBrandFactory("luxury")
	shoe = factory.CreateShoe()
	bag = factory.CreateBag()
	printDetails(shoe, bag)
}

func printDetails(s IShoe, b IBag) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Color: %s", s.GetColor())
	fmt.Println()
	fmt.Printf("Logo: %s", b.GetLogo())
	fmt.Println()
	fmt.Printf("Model: %s", b.GetModel())
	fmt.Println()
}
