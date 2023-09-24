package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/marcoagpegoraro/marco_blog/controllers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
)

func init() {
	initializers.LoadEnvPackages()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	engine := django.New("./views", ".django")

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
	app.Get("/", controllers.PostIndex)

	//Start App
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
