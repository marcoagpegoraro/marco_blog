package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
)

//go:embed views
var viewsAsssets embed.FS

func init() {
	initializers.LoadEnvPackages()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
	initializers.ConnectToS3()
}

func main() {
	// Create a new engine
	engine := django.NewPathForwardingFileSystem(http.FS(viewsAsssets), "/views", ".django")

	// register functions
	engine.AddFunc("getFirstImageUrlFromString", helpers.GetFirstImageUrlFromString)
	engine.AddFunc("formatDate", helpers.FormatDate)
	engine.AddFunc("camelCaseToCapitalizeFirstWorldAndAddSpaces", helpers.CamelCaseToCapitalizeFirstWorldAndAddSpaces)

	//Setup app
	app := fiber.New(fiber.Config{
		Views:         engine,
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Blog do marco v1.0.0",
	})

	app.Use(logger.New())

	//Configure App
	app.Static("/", "./public")

	//Routes
	Routes(app)

	//Start App
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
