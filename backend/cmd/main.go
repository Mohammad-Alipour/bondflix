package main

import (
	"net/http"

	"bondflix/backend/pkg/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		ws.ServeWs(c.Writer, c.Request)
	})

	r.Run(":8080")
}
