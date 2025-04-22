🔍 Detailed Purpose of Each Folder and File
broker/
Handles the core backend logic (written in Go):

main.go: Entry point of the broker application. Listens for connections.

broker.go: Handles socket connections, routing messages.

topic_manager.go: Manages writing messages to topic files.

file_lock.go: Implements file locking to avoid race conditions.

utils/logger.go: Optional logging helper for debugging.

producer_frontend/
Light HTML/JS interface for message producers:

index.html: Simple form UI for submitting topic & message.

script.js: Handles form submission and sends data to the Broker via sockets/WebSocket.

styles.css: Basic styling for layout.

consumer_frontend/
HTML/JS interface for message consumers:

index.html: UI for selecting a topic and offset (start/latest).

script.js: Connects to the broker and displays messages in real time or by offset.

styles.css: Page styling.

communication_protocol/
message.go: Shared Go struct for standardizing messages exchanged between components. Useful if Producer, Consumer, or test clients are CLI-based and written in Go.

test/
Go test files:

test_producer.go: Simulates a producer client sending messages.

test_consumer.go: Simulates a consumer fetching messages by offset.

Root Files
README.md: Describe how the system works and how to run it.

go.mod: Go module file for managing dependencies.