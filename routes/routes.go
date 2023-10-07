package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/marcoagpegoraro/marco_blog/controllers"
)

// Cheat Sheet - How to write routes:
// GET  /books                -- html table listing all books (alternatively /books/list to go with the DAREL acronym)
// GET  /books/add            -- display a form for adding a new book
// POST /books/add            -- adds a new book and redirects to /book/1 (where 1 is a new book id)
// GET  /books/1               -- display book 1 info (e.g. a customer view)
// GET  /books/1/edit          -- display a form to edit /book/1
// POST /books/1/edit          -- updates /book/1 and redirects to /book/1
// GET  /books/1/remove        -- maybe/probably optional
// POST /books/1/remove        -- normally /book/1/edit will have a delete button that handles "are you sure..?" and posts here, redirects to /books

func Routes(app *fiber.App) {
	//redirect rules
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/posts": "/",
		},
		StatusCode: 301,
	}))

	app.Get("/", controllers.GetIndex)

	app.Get("/login.php", controllers.GetLogin)
	app.Post("/login.php", controllers.PostLogin)

	app.Get("/posts/add", controllers.GetPostIndex)
	app.Post("/posts/add", controllers.PostPostIndex)

	app.Get("/posts/:id", controllers.GetOnePostIndex)
	app.Get("/posts/:id/edit", controllers.GetPostUpdate)
	app.Post("/posts/:id/edit", controllers.PostPostUpdate)
}
