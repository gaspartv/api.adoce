package router

import (
	"gaspartv/api.adoce/src/configs"
	"gaspartv/api.adoce/src/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Initialize(db *gorm.DB, env *configs.Env) error {
	router := gin.Default()

	authRoutes(router, db, env)
	userRoutes(router, db, env)

	if err := router.Run("0.0.0.0:" + env.Port); err != nil {
		return err
	}
	return nil
}

func authRoutes(router *gin.Engine, db *gorm.DB, env *configs.Env) {
	authService := service.NewAuthService(db, env)

	authGroup := router.Group("")
	{
		authGroup.POST("/sign-in", authService.Login)
	}
}

func userRoutes(router *gin.Engine, db *gorm.DB, env *configs.Env) {
	userService := service.NewUserService(db, env)

	userGroup := router.Group("users")
	{
		userGroup.POST("/create", userService.CreateUser)
	}
}
