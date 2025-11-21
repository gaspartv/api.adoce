package main

import (
	"gaspartv/api.adoce/src/configs"
	"gaspartv/api.adoce/src/internal/database"
	"gaspartv/api.adoce/src/internal/router"
)

func main() {
	env, err := configs.ValidateEnv()
	if err != nil {
		panic(err)
	}

	db, err := database.Initialize(env)
	if err != nil {
		panic(err)
	}

	if err := router.Initialize(db, env); err != nil {
		panic(err)
	}
}
