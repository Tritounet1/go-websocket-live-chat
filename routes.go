package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func helloworld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func LoginRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func RegisterRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func GetRoom(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func GetRooms(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func CreateRoom(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func DeleteRoom(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func SendMessageRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func WebSocketRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
	/*
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
			time.Sleep(time.Second)
		}
	*/
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "Bearer valid-token" {
			c.AbortWithStatusJSON(801, gin.H{"message": "Unauthorized"})
			return
		}

		c.Next()
	}
}
