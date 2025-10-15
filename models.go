package main

import (
	"context"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type Room struct {
	gorm.Model
	name     string
	password *string
	public   *bool
	owner    User
	users    []User
	messages []Message
}

type Image struct {
	gorm.Model
	name      string
	extension string
}

type Message struct {
	gorm.Model
	text   string
	images *[]Image
	sender User
	room   Room
}

type CacheMessage struct {
	gorm.Model
	user    User
	message Message
	isRead  bool
}

type Token struct {
	gorm.Model
	Token    string
	UserID   uint
	User     User
	ExpireAt time.Time
}

func StartDB() (*gorm.DB, context.Context) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate all schemas
	db.AutoMigrate(&User{}, &Room{}, &Image{}, &Message{}, &CacheMessage{}, &Token{})

	return db, ctx
}
