package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
)

var LogoffController = LogoffControllerStruct{}

type LogoffControllerStruct struct {
}

func (controller LogoffControllerStruct) Get(c *fiber.Ctx) error {
	c.ClearCookie("token")

	return c.RedirectToRoute("", fiber.Map{
		"messages": []dto.MessageDto{{Message: "Log off with success!", Type: "success"}},
	})
}
