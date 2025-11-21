package database

import (
	"gaspartv/api.adoce/src/configs"
	"gaspartv/api.adoce/src/internal/database/schemas"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func Initialize(env *configs.Env) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", env.DatabaseURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.AutoMigrate(&schemas.User{})

	return db, nil
}
