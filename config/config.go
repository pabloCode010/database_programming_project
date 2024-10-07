package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

var JwtKey string

func Config() {
	godotenv.Load(".env")

	JwtKey = os.Getenv("JWT_KEY")
}
