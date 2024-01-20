package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() Envs {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs := Envs{
		PORT:   os.Getenv("PORT"),
		DB_URL: os.Getenv("DB_URL"),
	}

	return envs
}
