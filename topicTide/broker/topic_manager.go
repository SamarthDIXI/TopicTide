package broker

import (
	"fmt"
	"log"
)

// HandleMessage processes the message and forwards it to file_lock for storage
func HandleMessage(filePath string, content []byte) error {
	
	// Log the incoming message and perform any necessary parsing/validation
	fmt.Printf("Handling message: %s\n", content)

	// Forward the message to file_lock for storage
	if err := WriteToFile(filePath , content); err != nil {
		log.Println("Error writing to file:", err)
		return err
	}

	return nil
}
