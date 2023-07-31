package user

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"ID"`
	FirstName string `gorm:"not null; size:50" json:"firstName"`
	LastName  string `gorm:"not null; size:50" json:"lastName"`
	Email     string `gorm:"not null; size:50; unique" json:"email"`
	Password  string `gorm:"not null; size:100" json:"-"`
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashPass)
	err = db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(u),
	})
}
