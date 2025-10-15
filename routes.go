package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func helloworld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! Time : %s", time.Now().Format(time.RFC3339Nano))
}

func LoginRoute(c *gin.Context) {
	var user User

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var password string = user.Password

	user, err = gorm.G[User](db).Where("username = ?", user.Username).First(ctx)

	if user.Password != password {
		c.String(http.StatusBadRequest, "Incorrect password")
		return
	}

	if err != nil {
		c.String(http.StatusBadRequest, "No user found")
		return
	} else {
		tokenExist, err := gorm.G[Token](db).Where("user_id = ?", user.ID).First(ctx)

		var token string

		if err != nil {
			token, err = GenerateToken(32)

			if err != nil {
				c.String(http.StatusBadRequest, "Error encounter while generating new token")
				return
			}

			new_token := Token{
				Token:    token,
				UserID:   user.ID,
				ExpireAt: time.Now().Add(7 * 24 * time.Hour),
			}

			err = gorm.G[Token](db).Create(ctx, &new_token)

			if err != nil {
				c.String(http.StatusBadRequest, "Error encounter while creating new token")
				return
			}
		} else {
			token = tokenExist.Token
		}

		c.JSON(http.StatusCreated, token)
	}
}

func RegisterRoute(c *gin.Context) {
	var user User

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = gorm.G[User](db).Where("username = ?", user.Username).First(ctx)

	if err == nil {
		c.String(http.StatusBadRequest, "A user with this username already exist")
		return
	}

	err = gorm.G[User](db).Create(ctx, &user)

	if err != nil {
		c.String(http.StatusBadRequest, "Error while creating a new user")
	} else {
		token, err := GenerateToken(32)

		if err != nil {
			c.String(http.StatusBadRequest, "Error encounter while generating new token")
			return
		}

		new_token := Token{
			Token:    token,
			UserID:   user.ID,
			ExpireAt: time.Now().Add(7 * 24 * time.Hour),
		}

		err = gorm.G[Token](db).Create(ctx, &new_token)

		if err != nil {
			c.String(http.StatusBadRequest, "Error encounter while creating new token")
			return
		}

		c.JSON(http.StatusCreated, token)
	}
}

func GetRoom(c *gin.Context) {
	var id int

	err := c.ShouldBind(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := gorm.G[Room](db).Where("id = ?", id).First(ctx)

	if err != nil {
		c.String(http.StatusBadRequest, "Error while searching the room")

	} else {
		c.JSON(http.StatusAccepted, room)
	}
}

func GetRooms(c *gin.Context) {
	rooms, err := gorm.G[Room](db).Find(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, &rooms)
	} else {
		c.JSON(http.StatusOK, &rooms)
	}
}

func CreateRoom(c *gin.Context) {
	var room Room

	err := c.ShouldBind(&room)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = gorm.G[Room](db).Create(ctx, &room)

	if err != nil {
		c.String(http.StatusBadRequest, "Error while creating a new room")

	} else {
		c.String(http.StatusCreated, "New room create")
	}
}

func DeleteRoom(c *gin.Context) {
	var id int

	err := c.ShouldBind(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = gorm.G[Room](db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		c.String(http.StatusBadRequest, "Error while deleting the room")

	} else {
		c.String(http.StatusCreated, "Room delete")
	}
}

func SendMessageRoute(c *gin.Context) {
	var message Message

	err := c.ShouldBind(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = gorm.G[Message](db).Create(ctx, &message)

	if err != nil {
		c.String(http.StatusBadRequest, "Error while creating a new room")

	} else {
		c.String(http.StatusCreated, "New room create")
	}
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
