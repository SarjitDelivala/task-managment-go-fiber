package routes

import (
	"github.com/SarjitDelivala/task-managment-go-fiber/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoadAllRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api", middlewares.ApiMiddleware)

	v1 := api.Group("/v1", middlewares.V1Middleware)
	LoadTaskRoutes(v1, db)
	LoadUsersRoutes(v1, db)
}
