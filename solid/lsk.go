package solid

import "fmt"

// Liskov Substitution Principle
//
// The Liskov Substitution Principle (LSP) states that if S is a subtype of T,
// then objects of type T may be replaced with objects of type S (i.e., an
// object of type T may be substituted with any object of a subtype S) without
// altering any of the desirable properties of the program (correctness,
// task performed, etc.).

type Sized interface {
	GetWidth() int
	GetHeight() int
	SetWidth(width int)
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	w := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * w
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected an area of ", expectedArea, ", but got ", actualArea)
}

type Square struct {
	Rectangle
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// Notice that the Square class inherits from the Rectangle class, and
// in order to enforce a square, it sets both width and height to the
// same value. However, notice that, when we use the UseIt function,
// we pass a Square object, which is a Sized object, and we expect
// the area to be 100, but it is not we receive 50. This is because the Square class
// set width in setHeight and height in setWidth, which breaks the logic of the UseIt.
// This is a violation of the Liskov Substitution Principle, because we expect the
// Square class to be a subtype of the Rectangle class, but it is not.
func mainLSK() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := &Square{}
	sq.SetWidth(5)
	UseIt(sq)
}
