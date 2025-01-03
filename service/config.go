package service

import (
	"os"
	"strconv"

	"github.com/LiteyukiStudio/go-logger/log"
	"github.com/joho/godotenv"
)

var (
	Token   = "server-status-be-q2dw9adh8"
	Timeout = 30
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

	if timeout, err := strconv.Atoi(os.Getenv("TIMEOUT")); err == nil {
		Timeout = timeout
	}
}
