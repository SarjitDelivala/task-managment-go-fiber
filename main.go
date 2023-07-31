package main

import (
	"fmt"
	"github.com/SarjitDelivala/task-managment-go-fiber/database"
	"github.com/SarjitDelivala/task-managment-go-fiber/routes/v1"
	"github.com/SarjitDelivala/task-managment-go-fiber/task"
	"github.com/SarjitDelivala/task-managment-go-fiber/user"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func main() {

	app := fiber.New()

	app.Get("/", helloWorld)

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Can't Load ENV file")
		return
	}

	db, err := database.Connect()
	if err != nil {
		fmt.Println("DB Connection Error")
		return
	}

	err = db.AutoMigrate(&user.User{}, &task.Task{})
	if err != nil {
		fmt.Println("DB Migration Error")
		return
	}

	routes.LoadAllRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
