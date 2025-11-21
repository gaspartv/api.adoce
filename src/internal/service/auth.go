package service

import (
	"gaspartv/api.adoce/src/configs"
	"gaspartv/api.adoce/src/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TokenMessage struct {
	Token string `json:"token"`
}

type AuthService struct {
	db  *gorm.DB
	env *configs.Env
}

func NewAuthService(db *gorm.DB, env *configs.Env) *AuthService {
	return &AuthService{
		db:  db,
		env: env,
	}
}

func (a *AuthService) Login(ctx *gin.Context) {
	var auth entity.Auth
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	println(auth.Email)
	println(auth.Password)

	response := TokenMessage{
		Token: "tokenString",
	}

	ctx.JSON(http.StatusOK, response)
}
