package entity

import (
	"strconv"
	"time"

	"gaspartv/api.adoce/src/configs"
	"gaspartv/api.adoce/src/internal/util"

	"github.com/google/uuid"
)

type User struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Name      string     `json:"name" gorm:"size:255;not null"`
	Email     string     `json:"email" gorm:"unique;size:255;not null"`
	Password  string     `json:"password" gorm:"size:255;not null"`
}

type CreateUserDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ResponseUserDto struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

func NewUser(user CreateUserDto, env *configs.Env) (*User, error) {
	bcrypt := util.Bcrypt{}

	salt, err := strconv.Atoi(env.BcryptSalt)
	if err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.Hash(user.Password, salt)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.New().String(),
		Email:    user.Email,
		Password: passwordHash,
	}, nil
}

func ResponseUser(user *User) ResponseUserDto {
	return ResponseUserDto{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}
}
