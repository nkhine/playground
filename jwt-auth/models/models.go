// models/models.go

package models

import (
	"go-tutorials/jwt-auth/database"

	"github.com/hashicorp/vault/builtin/logical/database"
	"gorm.io/gorm"
)

// User defines the user in the DB
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// CreateUserRecord creates a user record in the database
func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
