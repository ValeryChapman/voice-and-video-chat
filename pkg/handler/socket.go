package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// use default options
var upgrader = websocket.Upgrader{}

// routes -> for voice chat
var voice_routes = make(map[string][]*websocket.Conn)

// socket -> for voice chat
func voiceSocket(w http.ResponseWriter, r *http.Request, c *gin.Context) {

	// upgrade our raw HTTP connection to a websocket based one
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	// get room id
	roomId := c.Param("id")

	// add clients
	voice_routes[roomId] = append(voice_routes[roomId], ws)

	// start event loop
	for {

		// read message
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}

		// send message
		for client := range voice_routes[roomId] {

			client := voice_routes[roomId][client]

			// check users to send a one-way message
			if client != ws {

				// send
				err := client.WriteMessage(messageType, message)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					// ...
				}

			}
		}
	}
}

// routes -> for video chat
var video_routes = make(map[string][]*websocket.Conn)

// socket -> for video chat
func videoSocket(w http.ResponseWriter, r *http.Request, c *gin.Context) {

	// upgrade our raw HTTP connection to a websocket based one
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	// get room id
	roomId := c.Param("id")

	// add clients
	video_routes[roomId] = append(video_routes[roomId], ws)

	// start event loop
	for {

		// read message
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}

		// send message
		for client := range video_routes[roomId] {

			client := video_routes[roomId][client]

			// check users to send a one-way message
			if client != ws {

				// send
				err := client.WriteMessage(messageType, message)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					// ...
				}

			}
		}
	}
}
