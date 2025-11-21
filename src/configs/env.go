package configs

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Env struct {
	Port        string `validate:"required"`
	DatabaseURL string `validate:"required"`
	BcryptSalt  string `validate:"required"`
}

func ValidateEnv() (*Env, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var env Env
	env.Port = os.Getenv("PORT")
	env.DatabaseURL = os.Getenv("DATABASE_URL")
	env.BcryptSalt = os.Getenv("BCRYPT_SALT")

	validate := validator.New()
	if err := validate.Struct(env); err != nil {
		return nil, err
	}

	return &env, nil
}
