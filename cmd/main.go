package main

import (
	"github.com/fvdime/go-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectToDB()
	app := fiber.New()
	
	setupRoutes(app)
	app.Listen(":3000")
}
