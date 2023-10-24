package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/helpers"
)

var LoginController = LoginControllerStruct{}

type LoginControllerStruct struct {
}

func (controller LoginControllerStruct) Get(c *fiber.Ctx) error {
	return c.Render("pages/login/index", fiber.Map{
		"title": "Login",
	}, "layouts/login")
}

func (controller LoginControllerStruct) Post(c *fiber.Ctx) error {
	loginRequest := new(dto.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	usernameHash := os.Getenv("USERNAME_HASH")
	passwordHash := os.Getenv("PASSWORD_HASH")

	isUsernameCorrect := helpers.CompareHashStr(usernameHash, loginRequest.Username)
	isPasswordCorrect := helpers.CompareHashStr(passwordHash, loginRequest.Password)

	if !isUsernameCorrect || !isPasswordCorrect {
		return c.RedirectToRoute("", fiber.Map{
			"messages": []dto.MessageDto{{Message: "Username or password incorrect", Type: "warning"}},
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"Name":  loginRequest.Username,
		"Admin": true,
		"Exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.RedirectToRoute("", fiber.Map{
			"messages": []dto.MessageDto{{Message: "Error creating token", Type: "danger"}},
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		HTTPOnly: true,
		SameSite: "lax",
	})

	return c.RedirectToRoute("", fiber.Map{
		"messages": []dto.MessageDto{{Message: "Welcome!", Type: "success"}},
	})
}
