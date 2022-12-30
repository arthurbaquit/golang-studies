package main


type square struct {
	side float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) float64{
	return s.area()
}

func main(){
	square := square{
		side: 1,
	}
	circle := circle{
		radius: 1,
	}
	println(info(square))
	println(info(circle))
}