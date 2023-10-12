package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/marcoagpegoraro/marco_blog/dto"
)

func IsAuthenticated(c *fiber.Ctx) error {
	c.Locals("is_auth", false)
	jwtToken := c.Cookies("token", "")

	if jwtToken == "" {
		return c.Next()
	}

	var userClaim dto.UserClaim

	token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Next()
	}

	c.Locals("is_auth", true)
	return c.Next()
}
