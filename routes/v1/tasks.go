package routes

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/SarjitDelivala/task-managment-go-fiber/task"
)

func LoadTaskRoutes(r fiber.Router, db *gorm.DB) {

	r.Get("/tasks", func(c *fiber.Ctx) error {
		tasks, err := task.ListTasks(db)
		if err != nil {
			return err
		}
		return c.JSON(tasks)
	})
	r.Get("/tasks/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		t, err := task.GetTask(db, id)
		if err != nil {
			return err
		}
		return c.JSON(t)
	})
	r.Post("/tasks/", func(c *fiber.Ctx) error {
		t := new(task.Task)
		if err := c.BodyParser(t); err == nil {
			return err
		}
		t, err := task.CreateTask(db, t)
		if err != nil {
			return err
		}
		return c.JSON(t)
	})
	r.Patch("/tasks/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		t := new(task.Task)
		if err := c.BodyParser(t); err == nil {
			return err
		}
		t, err := task.UpdateTask(db, id, t)
		if err != nil {
			return err
		}
		return c.JSON(t)
	})
	r.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		err := task.DeleteTask(db, id)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Successfully Deleted",
		})
	})
	fmt.Println("V1 Routes Added")
}
