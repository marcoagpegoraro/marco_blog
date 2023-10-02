package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.GetIndex)

	app.Get("/posts", controllers.GetPostIndex)
	app.Get("/posts/:id", controllers.GetOnePostIndex)
	app.Post("/posts", controllers.PostPostIndex)
}
