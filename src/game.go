package main

import (
	"time"
)

type Game struct {
	Width    int
	Height   int
	Ball     Ball
	TickTime int
	Running  bool
}

type Ball struct {
	Pos       Position
	Direction int
	DeltaX    int
	DeltaY    int
}

type Position struct {
	Top  int
	Left int
}

func NewGame() Game {
	width := 300
	height := 300

	g := Game{
		width,
		height,
		Ball{
			Position{width/2 - 5, height/2 - 5}, // TODO remove magic numbers
			1,
			2,
			2,
		},
		100,
		false,
	}

	return g
}

func (g *Game) Update() {
	if g.Ball.Pos.Left >= g.Width-30 {
		g.Ball.Direction = -1
	} else if g.Ball.Pos.Left <= 30 {
		g.Ball.Direction = 1
	}
	g.Ball.Pos.Left += (5 * g.Ball.Direction)
}

func (g *Game) startGameLoop() {
	ticker := time.NewTicker(time.Duration(g.TickTime) * time.Millisecond)
	for range ticker.C {
		if g.Running {
			g.Update()
		}
	}
}
