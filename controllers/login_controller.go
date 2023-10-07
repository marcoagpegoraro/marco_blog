package controllers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/helpers"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("pages/login/index", fiber.Map{
		"title": "Login",
	}, "layouts/login")
}

func PostLogin(c *fiber.Ctx) error {
	loginRequest := new(dto.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	isUsernameCorrect := helpers.CompareHashStr(username, loginRequest.Username)
	isPasswordCorrect := helpers.CompareHashStr(password, loginRequest.Password)

	if !isUsernameCorrect || !isPasswordCorrect {
		return c.Redirect("/")
	}

	return c.Render("pages/login/index", fiber.Map{
		"title": "Login",
	}, "layouts/login")
}
