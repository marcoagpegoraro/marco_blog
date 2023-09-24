package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/controllers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
)

func init() {
	initializers.LoadEnvPackages()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	//Setup app
	app := fiber.New()

	//Routes
	app.Get("/", controllers.PostIndex)

	//Start App
	app.Listen(":" + os.Getenv("PORT"))
}
