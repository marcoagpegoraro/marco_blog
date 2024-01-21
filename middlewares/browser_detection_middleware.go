package middlewares

import (
	"github.com/dineshgowda24/browser"
	"github.com/gofiber/fiber/v2"
)

func DetectBrowser(c *fiber.Ctx) error {
	c.Locals("layout_path_prefix", "newer_browsers/")

	b, err := browser.NewBrowser(c.GetReqHeaders()["User-Agent"])

	if err != nil {
		return c.Next()
	}

	if b.Name() == "Internet Explorer" {
		setOldBrowserTrue(c)
	}

	return c.Next()
}

func setOldBrowserTrue(c *fiber.Ctx) {
	c.Locals("layout_path_prefix", "legacy_browsers/")
}
