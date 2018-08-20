package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
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

func main() {
	mux := http.NewServeMux()
	log := log.New(os.Stdout, "web ", log.LstdFlags)
	game := NewGame()
	app := &app{mux, log, nil, &game}

	go app.game.startGameLoop()

	mux.HandleFunc("/static/", app.static)
	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/websocket", app.ServeWebSocket)

	http.ListenAndServe(":8080", mux)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *app) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/html/index.html")
}

func (a *app) static(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	http.ServeFile(w, r, fmt.Sprintf("src/js/%s", splitPath[len(splitPath)-1]))
}

func (a *app) ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		a.log.Println(err)
		return
	}
	// Keep track of the new client inside the app
	c := client{ws, 0}
	a.clients = append(a.clients, &c)

	// Write loop inside a different goroutine
	go c.writePump(a)
	// Read loop inside this thread
	c.readPump(a)

	// TODO remove client from list on socket close
}
