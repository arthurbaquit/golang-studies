package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var connections = make(map[*websocket.Conn]int)


func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":3000", nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	websocketConn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Fatal(err)
	}
	remoteAddr := websocketConn.RemoteAddr().(*net.TCPAddr)
	remotePort := remoteAddr.Port
	connections[websocketConn] = remotePort

	go handleConnection(websocketConn)
}

func handleConnection(conn *websocket.Conn) {
	

	// someone just connected, tell everybody
	fmt.Println("User", connections[conn], "just connected.")
	for c  := range connections {
		c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("User %d just connected.", connections[conn])))
	}

	for {
		// wait for message from client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// someone just sent a message tell everybody
		for c := range connections {
			c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("User %d says: %s", connections[conn], message)))
		}
	}
	discPort := connections[conn]
	// remove the connection from connections slice
	delete(connections, conn)
	// someone just disconnected, tell everybody
	fmt.Println("User", discPort, "just disconnected.")
	for c := range connections {
		c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("User %d just disconnected.", discPort)))
	}
	conn.Close()
}

// //client code 
// let ws = new WebSocket("ws://localhost:8080");
// ws.onmessage = message => console.log(`Received: ${message.data}`);
// ws.send("Hello! I'm client")