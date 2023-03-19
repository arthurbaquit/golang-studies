package model

import (
	"image/color"
	"math"
	"math/rand"
	constants "multithreading/boids/constants"
	"sync"
	"time"
)

var (
	Green    = color.RGBA{0x00, 0xff, 0x00, 0xff}
	Boids    = make([]*Boid, constants.BoidCount)
	BoidGrid = make([][]int, constants.Width+1)
	RWLock   = sync.RWMutex{}
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
	acc := b.calcAcceleration()
	RWLock.Lock()
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
	RWLock.Unlock()
}

func (b *Boid) calcAcceleration() *Vector2D {
	upper := Vector2D{X: b.Position.X + constants.ViewRadius, Y: b.Position.Y + constants.ViewRadius}
	lower := Vector2D{X: b.Position.X - constants.ViewRadius, Y: b.Position.Y - constants.ViewRadius}
	avgVel := &Vector2D{X: 0, Y: 0}
	avgPos := &Vector2D{X: 0, Y: 0}
	avgSeparation := &Vector2D{X: 0, Y: 0}
	count := 0.0
	acc := &Vector2D{X: b.borderBounce(b.Position.X, constants.Width), Y: b.borderBounce(b.Position.Y, constants.Height)}
	RWLock.RLock()
	for i := math.Max(0, lower.X); i < math.Min(constants.Width, upper.X); i++ {
		for j := math.Max(0, lower.Y); j < math.Min(constants.Height, upper.Y); j++ {
			if otherId := BoidGrid[int(i)][int(j)]; otherId != -1 && otherId != b.Id {
				if dist := b.Position.Distance(Boids[otherId].Position); dist < constants.ViewRadius {
					count++
					avgVel = avgVel.Add(Boids[otherId].Velocity)
					avgPos = avgPos.Add(Boids[otherId].Position)
					avgSeparation = avgSeparation.Add(b.Position.Sub(Boids[otherId].Position).DivideByScalar(dist))
				}
			}
		}
	}
	RWLock.RUnlock()
	if count > 0 {
		avgVel = avgVel.DivideByScalar(count).Sub(b.Velocity).MultByScalar(constants.AdjRate)
		avgPos = avgPos.DivideByScalar(count).Sub(b.Position).MultByScalar(constants.AdjRate)
		avgSeparation = avgSeparation.MultByScalar(constants.AdjRate)
		acc = acc.Add(avgVel).Add(avgPos).Add(avgSeparation)
	}
	return acc
}

func (b *Boid) borderBounce(pos, limit float64) float64 {
	if pos < constants.ViewRadius {
		return 1 / pos
	} else if pos > limit-constants.ViewRadius {
		return 1 / (pos - limit)
	}
	return 0
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
