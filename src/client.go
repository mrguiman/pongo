package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

const PING_PERIOD time.Duration = 1 * time.Second
const GAME_UPDATE_PERDIOD time.Duration = 200 * time.Millisecond
const WRITE_WAIT time.Duration = 500 * time.Millisecond

type client struct {
	ws       *websocket.Conn
	playerID int
}

type GameInitMessage struct {
	Type       string
	Game       *Game
	MyPlayerID int
}

type GameUpdateMessage struct {
	Type string
	Game *Game
}

func (c *client) readPump(a *app) {
	defer func() {
		a.log.Println("Closing a connection")
		c.ws.Close()
	}()

	// These ensure we correctly read the pong answer from the peer
	// within a limited timeframe
	c.ws.SetReadDeadline(time.Now().Add(2 * time.Second))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(2 * time.Second))
		return nil
	})

	for {
		_, p, err := c.ws.ReadMessage()
		if err != nil {
			a.log.Println(err)
			break
		}

		switch string(p) {
		case "READY":
			c.playerID, err = a.game.RegisterPlayer()
			if err != nil {
				a.log.Println(err)
				c.write(websocket.CloseMessage, []byte(err.Error()))
				continue
			}

			data, err := json.Marshal(GameInitMessage{"INIT", a.game, c.playerID})
			if err != nil {
				a.log.Println(err)
			}

			err = c.write(websocket.TextMessage, data)
		}

		if err != nil {
			a.log.Println(err)
		}

	}
}

func (c *client) writePump(a *app) {
	gameTicker := time.NewTicker(GAME_UPDATE_PERDIOD)
	pingTicker := time.NewTicker(PING_PERIOD)

	// If any issue arises when attempting to write to the client, we stop the connection
	defer func() {
		gameTicker.Stop()
		pingTicker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case <-gameTicker.C:
			data, err := json.Marshal(GameUpdateMessage{"UPDATE", a.game})
			if err != nil {
				a.log.Println(err)
			}

			if err := c.write(websocket.TextMessage, data); err != nil {
				a.log.Println("Error sending Update: ", err)
			}
		case <-pingTicker.C:
			// We write ping messages at regular intervals to assert that the client is still connected
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *client) write(mt int, message []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
	return c.ws.WriteMessage(mt, message)
}
