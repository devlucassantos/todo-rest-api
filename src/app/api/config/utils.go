package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func loadEnvFile() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Environment Variables File not found: ", err)
	}
}
