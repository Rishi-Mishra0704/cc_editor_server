package sockets

import (
	"cc_editor_server/model"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var rooms = make(map[string]map[*websocket.Conn]bool) // connected clients per room
var broadcast = make(chan model.File)                 // broadcast channel

// Configure the WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Get the room from the query parameters
	room := r.URL.Query().Get("room")
	if room == "" {
		log.Println("Room not specified")
		return
	}

	// Initialize the room if not exists
	if _, ok := rooms[room]; !ok {
		rooms[room] = make(map[*websocket.Conn]bool)
	}

	// Register new client in the room
	rooms[room][ws] = true

	for {
		var file model.File
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&file)
		if err != nil {
			log.Printf("error: %v", err)
			delete(rooms[room], ws)
			break
		}
		// Set the room for the file message
		file.Room = room
		// Send the newly received message to the broadcast channel
		broadcast <- file
	}
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected in the same room
		for client := range rooms[msg.Room] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(rooms[msg.Room], client)
			}
		}
	}
}
