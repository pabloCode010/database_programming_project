package config

import "github.com/lpernett/godotenv"

func Config() {
	godotenv.Load(".env")
}
