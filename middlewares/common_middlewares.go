package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/marcoagpegoraro/marco_blog/initializers"
)

func CommonMiddlewares(app *fiber.App) {
	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(csrf.New(csrf.Config{
		KeyLookup: "cookie:csrf_",
	}))
	app.Use(func(c *fiber.Ctx) error {
		initializers.Cache.DeleteExpired()
		if initializers.Cache.ItemCount() > 15 {
			initializers.Cache.Flush()
		}
		return c.Next()
	})
	app.Static("/", "./public")
}
