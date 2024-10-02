package service

import (
	"github.com/LiteyukiStudio/go-logger/log"
	"github.com/joho/godotenv"
	"os"
)

var (
	Token = "server-status-be-q2dw9adh8"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Info("Error loading .env file")
	}
	token := os.Getenv("TOKEN")
	if token != "" {
		Token = token
	}
	log.Info("Token: ", Token)
}
