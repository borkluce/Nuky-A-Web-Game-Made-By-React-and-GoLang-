package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket connection upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Connection lost:", err)
			break
		}
		fmt.Printf("Message taken: %s\n", msg)

		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}

/*
func main() {
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("WebSocket server working on 8080 port...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server cannot start:", err)
	}
} */
