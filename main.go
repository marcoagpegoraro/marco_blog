package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
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
		ServerHeader:  "Apache/2.4.1 (Unix)",
		AppName:       "Marco's Blog v1.0.0",
	})

	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(csrf.New(csrf.Config{
		KeyLookup: "cookie:csrf_",
	}))
	app.Static("/", "./public")

	app.Use(middlewares.IsAuthenticated)
	routes.ApiRoutes(app)
	routes.Routes(app)

	app.Use(middlewares.OnlyAuthenticated)
	routes.RestrictedRoutes(app)

	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("DSN"))
	fmt.Println(os.Getenv("AWS_ACCESS_KEY_ID"))
	fmt.Println(os.Getenv("AWS_SECRET_ACCESS_KEY"))
	fmt.Println(os.Getenv("JWT_SECRET"))
	fmt.Println(os.Getenv("USERNAME_HASH"))
	fmt.Println(os.Getenv("PASSWORD_HASH"))
	fmt.Println(os.Getenv("ENV"))

	//Start App
	if os.Getenv("ENV") == "PROD" {
		log.Fatal(app.Listen("0.0.0.0:" + os.Getenv("PORT")))
	} else {
		log.Fatal(app.Listen(":" + os.Getenv("PORT")))
	}
}
