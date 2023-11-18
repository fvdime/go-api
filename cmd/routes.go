package main

import (
	"github.com/fvdime/go-api/handler"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handler.Get)

	app.Post("/create", handler.Create)
}