package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	ctx context.Context
)

/*
On peut s'enregistrer / se login.

Quand on est connecté -> token interne grâce à la table Token

On peut voir les rooms ou s'est déjà enregistré (on peut se connecter a une route et on l'enregistre dans notre liste de rooms)

On peux donc envoyer un message dans une room

Connexion websocket permet d'informer des messages reçu dans une room.

Un user qui est connécté à l'app est donc informer directement en websocket mais si il est pas on met dans la table en cache le message
pour faire une notification quand il ce reconnecte.
*/

func main() {
	db, ctx = StartDB()

	router := gin.New()

	// Create
	/*
		err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})

		// Read
		product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)       // find product with integer primary key
		products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // find product with code D42

		// Update - update product's price to 200
		err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
		// Update - update multiple fields
		err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, map[string]interface{}{"Price": 200, "Code": "F42"})

		// Delete - delete product
		err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
	*/

	router.GET("/", helloworld)

	router.POST("/auth/login", LoginRoute)
	router.POST("/auth/register", RegisterRoute)

	router.POST("/api/ws", AuthMiddleware(), WebSocketRoute)

	router.GET("/api/room", AuthMiddleware(), GetRoom)
	router.GET("/api/rooms", AuthMiddleware(), GetRooms)
	router.POST("/api/room", AuthMiddleware(), CreateRoom)
	router.DELETE("/api/room", AuthMiddleware(), DeleteRoom)

	router.POST("/api/message", AuthMiddleware(), SendMessageRoute)

	router.Run("localhost:3000")
}
