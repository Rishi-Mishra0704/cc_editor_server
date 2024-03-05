package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
)

/*
1. Initialization:
   - A map called 'clients' is created to store WebSocket connections.
   - A channel called 'broadcast' is created to broadcast messages to all connected clients.

2. Struct Definition:
   - A struct named 'File' is defined to represent file data, including content and file extension.

3. Main Function:
   - Configures the WebSocket route '/ws'.
   - Starts listening for incoming connections.
   - Sets up Cross-Origin Resource Sharing (CORS) to allow requests from any origin.

4. handleConnections Function:
   - Accepts incoming HTTP requests and upgrades them to WebSocket connections.
   - Registers clients in the 'clients' map.
   - Reads incoming JSON messages representing files from clients.
   - Broadcasts the received file messages to all connected clients.

5. handleMessages Function:
   - Continuously listens for messages on the 'broadcast' channel.
   - Sends each received message to all connected clients.

6. WebSocket Upgrader Configuration:
   - Configures the WebSocket upgrader with a custom 'CheckOrigin' function allowing connections from any origin.

7. CORS Configuration:
   - Sets up Cross-Origin Resource Sharing (CORS) to allow requests from any origin, with specific allowed methods and headers.
*/

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan File)              // broadcast channel

// Configure the WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define the message structure
type File struct {
	Content       string `json:"content"`
	FileExtension string `json:"fileExtension"`
}

func main() {
	// Configure WebSocket route
	http.HandleFunc("/ws", handleConnections)
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)
	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("Server started on :8000")
	err := http.ListenAndServe(":8000", cors(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Register new client
	clients[ws] = true

	for {
		var file File
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&file)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- file
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
