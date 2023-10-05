package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.GetIndex)

	// GET  /books                -- html table listing all books (alternatively /books/list to go with the DAREL acronym)
	// GET  /books/add            -- display a form for adding a new book
	// POST /books/add            -- adds a new book and redirects to /book/1 (where 1 is a new book id)
	// GET  /books/1               -- display book 1 info (e.g. a customer view)
	// GET  /books/1/edit          -- display a form to edit /book/1
	// POST /books/1/edit          -- updates /book/1 and redirects to /book/1
	// GET  /books/1/remove        -- maybe/probably optional
	// POST /books/1/remove        -- normally /book/1/edit will have a delete button that handles "are you sure..?" and posts here, redirects to /books

	app.Get("/posts/add", controllers.GetPostIndex)
	app.Get("/posts/:id", controllers.GetOnePostIndex)
	app.Post("/posts", controllers.PostPostIndex)
}
