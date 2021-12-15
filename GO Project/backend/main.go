package main

import (
	"fmt"		//fmt import formats I/O with functions for printing and reading
	"net/http"	//this import is for request-response handling for a web application

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

// defining the WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)	//shows a resulting error in endpoint connection
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

//Naming server and sending inital HTTP set up to the server
func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {		//mapping websocket endpoint to serveWs function
		serveWs(pool, w, r)
	})
}

func main() {						
	fmt.Println("Distributed Chat App v0.01")	//simple print to show that it is a chat app
	setupRoutes()					//setupRoutes function call
	http.ListenAndServe(":8080", nil)		//listen on port 8080 while disregarding any secondary parameter (nil) for now
}
