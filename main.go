package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type app struct {
	mux     *http.ServeMux
	log     *log.Logger
	clients []*client
	game    *Game
}

type Message struct {
	Type string
	Game Game
}

func main() {
	mux := http.NewServeMux()
	log := log.New(os.Stdout, "web ", log.LstdFlags)
	game := NewGame()
	app := &app{mux, log, nil, &game}

	go app.game.startGameLoop()

	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/websocket", app.ServeWebSocket)

	http.ListenAndServe(":8080", mux)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *app) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func (a *app) ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		a.log.Println(err)
		return
	}
	// Keep track of the new client inside the app
	c := client{ws}
	a.clients = append(a.clients, &c)

	// Write loop inside a different goroutine
	go c.writePump(a)
	// Read loop inside this thread
	c.readPump(a)

	// TODO remove client from list on socket close
}
