package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/mapper"
)

func GetPostIndex(c *fiber.Ctx) error {
	return c.Render("posts/index", fiber.Map{
		"Title": "Create new post",
	}, "layouts/main")
}

func PostPostIndex(c *fiber.Ctx) error {
	post := new(dto.PostRequest)

	if err := c.BodyParser(post); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	postModel := mapper.MapPostRequestToPostModel(*post)

	fmt.Println(postModel)

	initializers.DB.Create(&postModel)

	return c.Render("posts/index", fiber.Map{
		"Title": "Create new post",
	}, "layouts/main")
}
