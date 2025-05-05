package broker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"topicTide/communication_protocol"
)

// upgrade http connection to websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Printf("Incoming request Origin: %s\n", r.Header.Get("Origin"))
		return true
	},
}

// sanitizeFilename removes characters not allowed in filenames from topic name
func sanitizeFilename(name string) string {
	fmt.Print(name);
	re := regexp.MustCompile(`[^\w\-.]`)
	return re.ReplaceAllString(name, "_")
}

// handling producer request
func HandleProducer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleProducer function called!")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("WebSocket connection established in handleProducer!")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		var msg communication_protocol.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("JSON Unmarshal Error:", err)
			continue
		}

		// Use sanitized topic name as file name
		sanitizedTopic := sanitizeFilename(msg.Topic)

		// Change data type of content
		content := []byte(msg.Content)

		filePath := filepath.Join("topicFiles", sanitizedTopic+".txt")


		// Create the file if it doesn't exist
		file, err := os.OpenFile(filePath, os.O_CREATE, 0644)
		if err != nil {
			log.Println("Error creating file:", err)
			continue
		}
		file.Close()

		// Call message handler with raw message 
		if err := HandleMessage(filePath, content); err != nil {
			log.Println("Error in handling message:", err)
			continue
		}
	}
	fmt.Println("Exiting handleProducer for this connection.")
}

func Broker() {
	http.HandleFunc("/producer", HandleProducer)
	fmt.Println("Broker is running at http://localhost:8080/producer")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
