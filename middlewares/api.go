package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ApiMiddleware(c *fiber.Ctx) error {
	fmt.Println("API Middleware")

	return c.Next()
}
