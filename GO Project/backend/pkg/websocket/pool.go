package websocket

import "fmt"

type Pool struct {
	Register   chan *Client     //sends out "New user joined"
	Unregister chan *Client     //notifies the chat when a user disconnects
	Clients    map[*Client]bool // can use this to indicate active/inacctive
	Broadcast  chan Message     // loops through clients and send a message  ~ actually sends the message
}

func NewPool() *Pool {					//used for registering, unregistering, client mapping, and message broadcasting
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:		//defining a register case in the pool of clients
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))	//printing the number of clients
			for client, _ := range pool.Clients {		//simple for loop within range of the number of clients
				fmt.Println(client)
				//when a new user is added, the below message is updated to the chat window
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-pool.Unregister:	//defining an unregister case in the pool of clients 
			delete(pool.Clients, client)	//function to delete a user from the pool
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))	//printing the reducted number of clients
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})	//below message is updated to the chat window
			}
			break
		case message := <-pool.Broadcast:	//defining a broadcast case in the pool of clients
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
