package model

import (
	"image/color"
	"math"
	"math/rand"
	constants "multithreading/boids/constants"
	"time"
)

var (
	Green    = color.RGBA{0x00, 0xff, 0x00, 0xff}
	Boids    = make([]*Boid, constants.BoidCount)
	BoidGrid = make([][]int, constants.Width+1)
)

type Boid struct {
	Position *Vector2D
	Velocity *Vector2D
	Id       int
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	// update velocity with acc
	acc := b.calcAccelaeration()
	b.Velocity = b.Velocity.Add(acc).Limit(1, -1)
	// update grid
	BoidGrid[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	BoidGrid[int(b.Position.X)][int(b.Position.Y)] = b.Id
	next := b.Position.Add(b.Velocity)
	if next.X >= constants.Width || next.X < 0 {
		b.Velocity = &Vector2D{Y: b.Velocity.Y, X: -b.Velocity.X}
	}
	if next.Y >= constants.Height || next.Y < 0 {
		b.Velocity = &Vector2D{Y: -b.Velocity.Y, X: b.Velocity.X}
	}
	// update grid
	// time.Sleep(5000 * time.Millisecond)

}

func (b *Boid) calcAccelaeration() *Vector2D {
	upper := Vector2D{X: b.Position.X + constants.ViewRadius, Y: b.Position.Y + constants.ViewRadius}
	lower := Vector2D{X: b.Position.X - constants.ViewRadius, Y: b.Position.Y - constants.ViewRadius}
	avgVel := &Vector2D{X: 0, Y: 0}
	count := 0.0
	for i := math.Max(0, lower.X); i < math.Min(constants.Width, upper.X); i++ {
		for j := math.Max(0, lower.Y); j < math.Min(constants.Height, upper.Y); j++ {
			if otherId := BoidGrid[int(i)][int(j)]; otherId != -1 && otherId != b.Id {
				if dist := b.Position.Distance(Boids[otherId].Position); dist < constants.ViewRadius {
					avgVel = avgVel.Add(Boids[otherId].Velocity)
					count++
				}
			}
		}
	}
	if count > 0 {
		avgVel = avgVel.DivideByScalar(count).Sub(b.Velocity).MultByScalar(constants.AdjRate)
	}
	return avgVel
}

func CreateNew(id int) {
	var b = new(Boid)
	b.Id = id
	b.Position = &Vector2D{X: rand.Float64() * constants.Width, Y: rand.Float64() * constants.Height}
	b.Velocity = &Vector2D{X: rand.Float64()*2 - 1, Y: rand.Float64()*2 - 1}
	// update grid
	BoidGrid[int(b.Position.X)][int(b.Position.Y)] = b.Id
	Boids[id] = b
	go b.start()
}
