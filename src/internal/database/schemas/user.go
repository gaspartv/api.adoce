package schemas

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id       uint `gorm:"primary_key"`
	Name     string
	Email    string `gorm:"unique_index"`
	Password string `gorm:"not null"`
}
