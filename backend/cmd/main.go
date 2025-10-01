package main

import (
	"log"
	"net/http"

	"github.com/Mohammad-Alipour/bondflix/backend/internal/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// s
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	r := gin.Default()
	hub := ws.NewHub()
	go hub.Run()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}
		client := &ws.Client{Conn: conn, Send: make(chan []byte)}
		hub.Register <- client

		go client.ReadPump(hub)
		go client.WritePump()
	})

	r.Run(":8080")
}
