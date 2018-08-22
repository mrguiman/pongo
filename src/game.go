package main

import (
	"errors"
	"math"
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
	Players  map[int]Player
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
	Width     int
	Height    int
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
			10,
			0,
			10,
			10,
		},
		100,
		false,
		make(map[int]Player),
	}

	return g
}

func (g *Game) Update() {
	if g.Ball.Direction == 1 && ballCollidesHorizontally(g.Ball, g.Players[2]) && ballCollidesVertically(g.Ball, g.Players[2]) {
		g.Ball.DeltaY = getDeltaY(g.Ball, g.Players[2])
		g.Ball.Direction = -1
	} else if g.Ball.Direction == -1 && ballCollidesHorizontally(g.Ball, g.Players[1]) && ballCollidesVertically(g.Ball, g.Players[1]) {
		g.Ball.DeltaY = getDeltaY(g.Ball, g.Players[1])
		g.Ball.Direction = 1
	}

	if g.ballCollidesWithFloorOrCeiling() {
		g.Ball.DeltaY = g.Ball.DeltaY * -1
	}

	g.Ball.Pos.X += (g.Ball.DeltaX * g.Ball.Direction)
	g.Ball.Pos.Y += g.Ball.DeltaY
}

func ballCollidesHorizontally(ball Ball, player Player) bool {
	if ball.Direction == 1 {
		return ball.Pos.X+ball.Width >= player.Pos.X
	} else {
		return ball.Pos.X <= player.Pos.X+player.Width
	}
}

func ballCollidesVertically(ball Ball, player Player) bool {
	playerCollisionY := [2]int{player.Pos.Y, player.Pos.Y + player.Height}

	if ball.Pos.Y+ball.Height < playerCollisionY[0] {
		return false
	} else if ball.Pos.Y > playerCollisionY[1] {
		return false
	} else {
		return true
	}
}

func getDeltaY(ball Ball, player Player) int {
	playerCenter := player.Pos.Y + player.Height/2
	ballCenter := ball.Pos.Y + ball.Height/2

	return int(math.Round(float64(ballCenter-playerCenter) / 3))
}

func (g *Game) startGameLoop() {
	ticker := time.NewTicker(time.Duration(g.TickTime) * time.Millisecond)
	for range ticker.C {
		if len(g.Players) == 2 && !g.Running {
			g.Running = true
		}

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
	g.Players[playerID] = newPlayer
	return newPlayer.ID, nil
}

func (g *Game) UnregisterPlayer(playerID int) {
	delete(g.Players, playerID)
}

func (g *Game) UpdatePlayerPosition(playerID int, newY int) {
	player := g.Players[playerID]
	player.Pos.Y = newY
	g.Players[playerID] = player
}

func (g *Game) ballCollidesWithFloorOrCeiling() bool {
	if g.Ball.DeltaY < 0 {
		return g.Ball.Pos.Y <= 0
	} else if g.Ball.DeltaY > 0 {
		return g.Ball.Pos.Y+g.Ball.Height >= g.Height
	} else {
		return false
	}
}
