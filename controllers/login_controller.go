package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/services"
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

	usernameHash := os.Getenv("USERNAME_HASH")
	passwordHash := os.Getenv("PASSWORD_HASH")

	isUsernameCorrect := helpers.CompareHashStr(usernameHash, loginRequest.Username)
	isPasswordCorrect := helpers.CompareHashStr(passwordHash, loginRequest.Password)

	posts := services.GetPosts(c)

	if !isUsernameCorrect || !isPasswordCorrect {
		return c.Render("pages/index/index", fiber.Map{
			"title":    "Home",
			"posts":    posts,
			"messages": []dto.MessageDto{{Message: "Username or password incorrect", Type: "warning"}},
		}, "layouts/main")
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  loginRequest.Username,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Render("pages/index/index", fiber.Map{
			"title":    "Home",
			"posts":    posts,
			"messages": []dto.MessageDto{{Message: "Error creating token", Type: "danger"}},
		}, "layouts/main")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		HTTPOnly: true,
		SameSite: "lax",
	})

	return c.Render("pages/index/index", fiber.Map{
		"title":    "Home",
		"posts":    posts,
		"messages": []dto.MessageDto{{Message: "Welcome!", Type: "success"}},
	}, "layouts/main")
}
