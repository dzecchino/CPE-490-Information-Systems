package websocket				//allows for browser-server communiction

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{		//defining an upgrader from standard http to websocket endpoint
	ReadBufferSize:  1024,			//necessary buffer sizes
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },	//validates request origin
}

//below function used to reply to client with an HTTP response error
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)		//logs and prints the error
		return nil, err
	}

	return conn, nil
}
