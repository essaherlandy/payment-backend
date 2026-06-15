package middleware

import "github.com/gofiber/fiber/v2"

func Role(roles ...string) fiber.Handler {

	return func(c *fiber.Ctx) error {

		userRole := c.Locals("role").(string)

		for _, role := range roles {

			if role == userRole {
				return c.Next()
			}
		}

		return c.Status(403).JSON(
			fiber.Map{
				"message": "access denied",
			},
		)
	}
}
