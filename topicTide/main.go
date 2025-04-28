package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Printf("Incoming request Origin: %s\n", r.Header.Get("Origin")) // Log the Origin header
		return true // Allow all origins 
	},
}

func handleProducer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleProducer function called!")
	fmt.Printf("Request Method: %s\n", r.Method)
	fmt.Printf("Request URL: %s\n", r.URL.Path)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("WebSocket connection established in handleProducer!")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}
		fmt.Printf("Received Message (Type %d): %s\n", messageType, message)
	}
	fmt.Println("Exiting handleProducer for this connection.")
}

func main() {
	http.HandleFunc("/", handleProducer)
	fmt.Println("Broker running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}