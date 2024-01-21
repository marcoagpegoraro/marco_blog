package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
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
	initializers.ConfigureCache()
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
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Apache/2.4.1 (Unix)",
		AppName:       "Marco's Blog v1.0.0",
	})

	middlewares.CommonMiddlewares(app)
	app.Use(middlewares.DetectBrowser)

	app.Use(middlewares.IsAuthenticated)
	routes.ApiRoutes(app)
	routes.Routes(app)

	app.Use(middlewares.OnlyAuthenticated)
	routes.RestrictedRoutes(app)

	//Start App
	log.Fatal(app.Listen("0.0.0.0:" + os.Getenv("PORT")))
}
