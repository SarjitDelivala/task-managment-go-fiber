package routes

import (
	"github.com/SarjitDelivala/task-managment-go-fiber/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoadUsersRoutes(r fiber.Router, db *gorm.DB) {
	r.Post("/users", func(c *fiber.Ctx) error {
		u := new(user.User)
		err := c.BodyParser(u)
		if err != nil {
			return err
		}
		u, err = user.CreateUser(db, u)
		if err != nil {
			return err
		}
		return c.JSON(u)
	})
}
