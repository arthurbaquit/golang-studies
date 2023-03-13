package main

import (
	constants "multithreading/boids/constants"
	model "multithreading/boids/model"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range model.Boids {
		screen.Set(int(b.Position.X+1), int(b.Position.Y), model.Green)
		screen.Set(int(b.Position.X-1), int(b.Position.Y), model.Green)
		screen.Set(int(b.Position.X), int(b.Position.Y+1), model.Green)
		screen.Set(int(b.Position.X), int(b.Position.Y-1), model.Green)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return constants.Width, constants.Height
}

func main() {
	// initialize grid with -1
	for i := 0; i < constants.Width+1; i++ {
		model.BoidGrid[i] = make([]int, constants.Height+1)
		for j := 0; j < constants.Height+1; j++ {
			model.BoidGrid[i][j] = -1
		}
	}
	for i := 0; i < constants.BoidCount; i++ {
		model.CreateNew(i)
	}
	ebiten.SetWindowSize(constants.Width, constants.Height)
	ebiten.SetWindowTitle("Boids")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
