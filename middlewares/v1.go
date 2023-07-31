package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func V1Middleware(c *fiber.Ctx) error {
	fmt.Println("V1 Middleware")
	return c.Next()
}
