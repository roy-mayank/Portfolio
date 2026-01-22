package main

// Functionalities to implement:
//    1. Websocker for real time comms for google docs like effect
//    2. Display active users based on hearbeats
import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },//CORS
}

type Hub struct {
	clients   map[*websocket.Conn]bool // Store of all active connections
	broadcast chan []byte // Multicast channel??
	mutex     sync.Mutex // temp low effort race condition prevention
}

func (h *Hub) run() {
	// Infinite loop to listen for messages 
	for {
		message := <-h.broadcast
		h.mutex.Lock()
		for client := range h.clients { // Loop through all clients available
			err := client.WriteMessage(websocket.TextMessage, message)
			// Client discconection handling
			if err != nil {
				client.Close()
				delete(h.clients, client)
			}
		}
		h.mutex.Unlock()
	}
}

func main() {
	hub := &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		hub.mutex.Lock()
		hub.clients[conn] = true
		hub.mutex.Unlock()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil { break }
			hub.broadcast <- msg
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}