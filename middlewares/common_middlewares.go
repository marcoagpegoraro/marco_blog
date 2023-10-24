package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CommonMiddlewares(app *fiber.App) {
	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(csrf.New(csrf.Config{
		KeyLookup: "cookie:csrf_",
	}))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Method() != "GET"
		},
		Expiration:   24 * time.Hour,
		CacheControl: true,
	}))
	app.Static("/", "./public")
}
