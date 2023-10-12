package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/middlewares"
	"github.com/marcoagpegoraro/marco_blog/routes"
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
		AppName:       "Marco's Blog v1.0.0",
	})

	app.Use(logger.New())
	app.Use(cors.New())
	app.Static("/", "./public")

	app.Use(middlewares.IsAuthenticated)
	routes.Routes(app)

	app.Use(middlewares.OnlyAuthenticated)
	routes.RestrictedRoutes(app)

	//Start App
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
