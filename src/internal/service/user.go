package service

import (
	"fmt"
	"gaspartv/api.adoce/src/configs"
	"gaspartv/api.adoce/src/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db  *gorm.DB
	env *configs.Env
}

func NewUserService(db *gorm.DB, env *configs.Env) *UserService {
	return &UserService{
		db:  db,
		env: env,
	}
}

func (s *UserService) CreateUser(ctx *gin.Context) {
	var createUserDto entity.CreateUserDto
	if err := ctx.ShouldBindJSON(&createUserDto); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	exists := s.UserExists(createUserDto.Email)
	if exists {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	userCreate, err := entity.NewUser(entity.CreateUserDto{
		Name:     createUserDto.Name,
		Email:    createUserDto.Email,
		Password: createUserDto.Password,
	}, s.env)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Create(&userCreate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userResponse := entity.ResponseUser(userCreate)
	ctx.JSON(http.StatusCreated, userResponse)
}

func (s *UserService) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user entity.User
	if err := s.db.First(&user, &id).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	userResponse := entity.ResponseUser(&user)
	c.JSON(http.StatusOK, userResponse)
}

func (s *UserService) UserExists(email string) bool {
	var user entity.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

func (s *UserService) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
