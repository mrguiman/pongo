package main

import (
	"errors"
	"time"
)

const PLAYER_WIDTH int = 10
const PLAYER_HEIGHT int = 60

type Game struct {
	Width    int
	Height   int
	Ball     Ball
	TickTime int
	Running  bool
	Players  []Player
}

type Player struct {
	ID     int
	Width  int
	Height int
	Pos    Position
}

type Ball struct {
	Pos       Position
	Direction int
	DeltaX    int
	DeltaY    int
}

type Position struct {
	X int
	Y int
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
		[]Player{},
	}

	return g
}

func (g *Game) Update() {
	if g.Ball.Pos.X >= g.Width-30 {
		g.Ball.Direction = -1
	} else if g.Ball.Pos.X <= 30 {
		g.Ball.Direction = 1
	}
	g.Ball.Pos.X += (5 * g.Ball.Direction)
}

func (g *Game) startGameLoop() {
	ticker := time.NewTicker(time.Duration(g.TickTime) * time.Millisecond)
	for range ticker.C {
		if g.Running {
			g.Update()
		}
	}
}

func (g *Game) RegisterPlayer() (int, error) {
	if len(g.Players) >= 2 {
		return 0, errors.New("There's already enough players connected")
	}

	var playerID, initialX int
	if len(g.Players) == 1 {
		playerID = 2
		initialX = g.Width - (PLAYER_WIDTH * 2)
	} else {
		playerID = 1
		// Position first player to the left
		initialX = PLAYER_WIDTH
	}

	initialY := (g.Height / 2) - (PLAYER_HEIGHT / 2)

	newPlayer := Player{playerID, PLAYER_WIDTH, PLAYER_HEIGHT, Position{initialX, initialY}}
	g.Players = append(g.Players, newPlayer)
	return newPlayer.ID, nil
}
