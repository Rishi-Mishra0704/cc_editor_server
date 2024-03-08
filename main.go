package main

import (
	"cc_editor_server/sockets"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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

func main() {
	// Configure WebSocket route
	http.HandleFunc("/ws", sockets.HandleConnections)
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)
	// Start listening for incoming chat messages
	go sockets.HandleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("Server started on :8000")
	err := http.ListenAndServe(":8000", cors(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
