package schemas

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id       string `gorm:"unique_index;not null"`
	Name     string
	Email    string `gorm:"unique_index"`
	Password string `gorm:"not null"`
}
