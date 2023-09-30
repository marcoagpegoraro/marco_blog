package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/marcoagpegoraro/marco_blog/controllers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
)

//go:embed views
var viewsAsssets embed.FS

func init() {
	initializers.LoadEnvPackages()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	// engine := django.New("./views", ".django")
	// Create a new engine
	engine := django.NewPathForwardingFileSystem(http.FS(viewsAsssets), "/views", ".django")

	//Setup app
	app := fiber.New(fiber.Config{
		Views:         engine,
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Blog do marco v1.0.0",
	})

	//Configure App
	app.Static("/", "./public")

	//Routes

	app.Get("/", controllers.GetIndex)

	app.Get("/posts", controllers.GetPostIndex)
	app.Post("/posts", controllers.PostPostIndex)

	//Start App
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
