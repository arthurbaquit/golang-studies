package model

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v *Vector2D) Add(v2 *Vector2D) *Vector2D {
	return &Vector2D{X: v.X + v2.X, Y: v.Y + v2.Y}
}
func (v *Vector2D) AddByScalar(scalar float64) *Vector2D {
	return &Vector2D{X: v.X + scalar, Y: v.Y + scalar}
}

func (v *Vector2D) DivideByScalar(scalar float64) *Vector2D {
	return &Vector2D{X: v.X / scalar, Y: v.Y / scalar}
}

func (v *Vector2D) Sub(v2 *Vector2D) *Vector2D {
	return &Vector2D{X: v.X - v2.X, Y: v.Y - v2.Y}
}
func (v *Vector2D) SubByScalar(scalar float64) *Vector2D {
	return &Vector2D{X: v.X - scalar, Y: v.Y - scalar}
}

func (v *Vector2D) MultByScalar(scalar float64) *Vector2D {
	return &Vector2D{X: v.X * scalar, Y: v.Y * scalar}
}

func (v *Vector2D) MultByVector(v2 *Vector2D) *Vector2D {
	return &Vector2D{X: v.X * v2.X, Y: v.Y * v2.Y}
}

func (v *Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2D) Limit(upper, lower float64) *Vector2D {
	return &Vector2D{X: math.Min(math.Max(v.X, lower), upper), Y: math.Min(math.Max(v.Y, lower), upper)}
}

func (v *Vector2D) Distance(v2 *Vector2D) float64 {
	return math.Sqrt(math.Pow(v2.X-v.X, 2) + math.Pow(v2.Y-v.Y, 2))
}
