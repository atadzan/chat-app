package main

import (
	"fmt"
	"github.com/atadzan/chat-app/backend/pkg/websocket"
	"log"
	"net/http"
)

// define websocket endpoint
func serverWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}
