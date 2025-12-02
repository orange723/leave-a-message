package main

import (
	"log"

	"leave-a-message/api"
	"leave-a-message/api/message"
	"leave-a-message/database"
	"leave-a-message/server"

	"github.com/joho/godotenv"
)

func main() {
	// not use
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Server initialization
	app := server.Create()

	// Migrations
	database.DB.AutoMigrate(&message.Message{})

	// Api routes
	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
