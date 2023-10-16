package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/controllers"
)

func RestrictedRoutes(app *fiber.App) {
	app.Get("/posts/add", controllers.GetPostIndex)
	app.Post("/posts/add", controllers.PostPostIndex)

	app.Get("/posts/:id/edit", controllers.GetPostUpdate)
	app.Post("/posts/:id/edit", controllers.PostPostUpdate)
}
