package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func GetIndex(c *fiber.Ctx) error {

	var posts []models.Post
	initializers.DB.Find(&posts)

	// fmt.Println(posts)

	for index, _ := range posts {
		fmt.Println(index)
		// fmt.Println(element)
	}

	return c.Render("index/index", fiber.Map{
		"Title": "Home",
		"Posts": posts,
	}, "layouts/main")
}
