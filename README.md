Go Chat Application

A simple real-time, one-to-one chat application built with Go and WebSockets. This project is a learning project focused on understanding Go through building a working application.

Current Stage: WebSocket Chat

Features implemented

Starts an HTTP server on port 8080.

Establishes WebSocket connections at /webSocket.

Identifies connected users through a user URL query parameter.

Registers and unregisters users in a shared chat hub.

Accepts JSON chat messages containing a recipient username and message text.

Routes messages to the intended recipient.

Detects disconnected connections and removes users from the hub.

Rejects connections that do not supply a username.

Technologies

Go — application logic and HTTP server.

Gorilla WebSocket — WebSocket connections and messaging.

JSON — the format used for incoming chat messages.

Project structure

go_chat/
├── main.go
└── server/
├── server.go
├── hub.go
└── chatMessage.go

main.go starts the HTTP server and creates the shared hub.

server/server.go handles WebSocket connections and incoming messages.

server/hub.go stores active user connections and provides registration and message-routing methods.

server/chatMessage.go defines the structure of an incoming chat message.

How it works

A client connects to /webSocket and supplies a username.

The server upgrades the HTTP connection to a WebSocket connection.

The hub stores the username and its connection.

The connected client sends a JSON message containing the recipient and text.

The server decodes the JSON and asks the hub to send the text to the recipient.

The recipient receives the message through their WebSocket connection.

When a connection closes, the server unregisters that user.

Running the application

Open the project in GoLand.

Run main.go.

Keep the server running while testing clients.

The server listens on port 8080.

Testing with GoLand

Create or use a chat.http file with the following requests:

WEBSOCKET ws://localhost:8080/webSocket?user=daniel

###

WEBSOCKET ws://localhost:8080/webSocket?user=ezra

Run each WebSocket request to establish two separate client connections.

From Daniel's connection, send:

{
"RecipientUsername": "ezra",
"Text": "Hi Ezra!"
}

Ezra should receive:

Hi Ezra!

The JSON property names match the exported fields in ChatMessage.

Current limitations

No authentication: clients can choose a username in the URL.

No database: messages are not stored permanently.

No offline messaging: a recipient must be connected to receive a message.

No conversations or user accounts: the current design is a basic username-to-username chat.

Concurrency safety needs improvement: the hub mutex protects the connection map, but it does not coordinate concurrent writes to a WebSocket connection.

Limited error handling: malformed or incomplete messages are currently ignored.

Next development stage

Before adding more chat features, revisit Go fundamentals and build a stronger understanding of the existing code.

Topics to learn include:

Structs, pointers, methods, and interfaces.

Functions, return values, errors, and defer.

Maps, slices, and package imports.

HTTP handlers and request/response objects.

JSON encoding and decoding.

Goroutines, channels, mutexes, and safe concurrent WebSocket writes.

Meaningful tests for message routing and connection lifecycle.

Project goal

Build a working one-to-one chat application while learning to write, explain, test, and debug Go code independently.

The current milestone is a functioning WebSocket chat prototype. Authentication, persistence, concurrency improvements, and production deployment remain future work.