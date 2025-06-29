package middlewares

import (
	"7solutions/backend/common/authorization"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AccessToken(c *fiber.Ctx) error {
	var accessToken string
	cookie := c.Cookies("Accesstoken")

	authorizationHeader := c.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else {
		accessToken = cookie
	}

	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"status":  false,
			"message": "unauthorized",
			"data":    "",
		})
	}

	jwtHS256 := authorization.NewJWT_HS256()

	sub := authorization.AppAuthorizationClaim{}
	err := jwtHS256.ValidateToken(accessToken, &sub)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"status":  false,
			"message": err.Error(),
			"data":    "",
		})
	}

	c.Locals("user_id", sub.UserId)

	return c.Next()
}
