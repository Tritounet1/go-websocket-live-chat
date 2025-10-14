package main

import (
	"net/http"
	"tidy/tests"
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
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		time.Sleep(time.Second)
	}
}

/*
user:
	id,
	username,
	password

room:
	id,
	name,
	password?, : string
	public? : bool
	owner: user
	users: user[]
	messages: message[]

message:
	id,
	text,
	images?, : string (liens d'images)
	user: user
	room: room

cache:
	id,
	user: user
	message: message
	isRead: bool

On peut s'enregistrer / se login.

Quand on est connécté -> token JWT.

On peut voir les rooms ou s'est déjà enregistré (on peut se connecter a une route et on l'enregistre dans notre liste de rooms)

On peux donc envoyer un message dans une room

Connexion websocket permet d'informer des messages reçu dans une room.

Un user qui est connécté à l'app est donc informer directement en websocket mais si il est pas on met dans la table en cache le message
pour faire une notification quand il ce reconnecte.
*/

func main() {
	tests.Test_sum()

	router := gin.New()

	router.GET("/", helloworld)

	router.POST("/auth/login", LoginRoute)
	router.POST("/auth/register", RegisterRoute)

	//

	router.GET("/ws", WebSocketRoute)

	router.GET("/api/room", GetRoom)
	router.GET("/api/rooms", GetRooms)
	router.POST("/api/room", CreateRoom)
	router.DELETE("/api/room", DeleteRoom)

	router.POST("/api/message", SendMessageRoute)

	router.Run("localhost:3000")
}
