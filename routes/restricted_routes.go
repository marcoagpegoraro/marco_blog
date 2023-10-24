package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/controllers"
)

func RestrictedRoutes(app *fiber.App) {
	app.Get("/posts/add", controllers.PostController.Get)
	app.Post("/posts/add", controllers.PostController.Post)

	app.Get("/posts/:id/edit", controllers.PostController.GetEditPost)
	app.Post("/posts/:id/edit", controllers.PostController.PostEditPost)
}
