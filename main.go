package main

import (
	"fmt"
	"go_chat/server"
	"log"
	"net/http"
)

func main() {

	fmt.Println("server starting")
	hub := server.NewHub()
	http.HandleFunc("/webSocket", server.NewChatHandler(hub))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
