package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_USER          string
	DB_PASS          string
	DB_PORT          string
	DB_HOST          string
	DB_NAME          string
	JWT_SECRET_TOKEN string
	DB_TYPE          string
}

func SetEnv() (*Env, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	env := &Env{
		DB_USER:          os.Getenv("DB_USER"),
		DB_PASS:          os.Getenv("DB_PASS"),
		DB_PORT:          os.Getenv("DB_PORT"),
		DB_HOST:          os.Getenv("DB_HOST"),
		DB_NAME:          os.Getenv("DB_NAME"),
		JWT_SECRET_TOKEN: os.Getenv("JWT_SECRET_TOKEN"),
		DB_TYPE:          os.Getenv("DB_TYPE"),
	}
	return env, nil
}

func GetSecret() ([]byte, error) {
	env, err := SetEnv()
	if err != nil {
		return nil, err
	}
	return []byte(env.JWT_SECRET_TOKEN), nil
}
