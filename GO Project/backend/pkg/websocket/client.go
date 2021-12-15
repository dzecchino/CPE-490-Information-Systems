package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {			//Client information structure
	ID   string			//Identifying client IDs as type string
	Conn *websocket.Conn		//This is a pointer to a websocket.Conn object
					//Conn is a generic stream-oriented network connection
	Pool *Pool			//pointer to the pool in which a client is a part of 
}

type Message struct {			//Message structure
	Type int    `json:"type"`
	Body string `json:"body"`	//Defining the body of the message as a string
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}	//Setting the message type and body information
		c.Pool.Broadcast <- message				//Looping through all clients in pool and sending through socket connection
		fmt.Printf("Message Received: %+v\n", message)
	}
}
