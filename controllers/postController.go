package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PostIndex(c *fiber.Ctx) error {
	return c.Render("posts/index", fiber.Map{
		"Title": "Create new post",
	}, "layouts/main")
}

func PostIndexPost(c *fiber.Ctx) error {
	type Post struct {
		PostTitle    string `json:"PostTitle" xml:"PostTitle" form:"PostTitle"`
		PostSubtitle string `json:"PostSubtitle" xml:"PostSubtitle" form:"PostSubtitle"`
		PostBody     string `json:"PostBody" xml:"PostBody" form:"PostBody"`
	}

	post := new(Post)

	if err := c.BodyParser(post); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	return c.Render("posts/index", fiber.Map{
		"Title": "Create new post",
	}, "layouts/main")
}
