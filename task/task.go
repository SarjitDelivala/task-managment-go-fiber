package task

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID          uint       `gorm:"primaryKey"`
	Title       string     `gorm:"not null"`
	DueDate     *time.Time `gorm:"null"`
	Description *string    `gorm:"null"`
	UserID      uint       `gorm:"not null"`
}

func ListTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Model(&Task{}).Find(&tasks).Error
	return tasks, err
}

func GetTask(db *gorm.DB, id int) (Task, error) {
	var task Task
	err := db.Model(&Task{}).Find(&task, id).Error
	return task, err
}

func CreateTask(db *gorm.DB, task *Task) (*Task, error) {
	err := db.Create(&task).Error
	return task, err
}

func UpdateTask(db *gorm.DB, id int, task *Task) (*Task, error) {
	err := db.Updates(&task).Error
	return task, err
}

func DeleteTask(db *gorm.DB, id int) error {
	err := db.Delete(&Task{}, id).Error
	return err
}
