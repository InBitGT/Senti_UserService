package config

import (
	"github.com/joho/godotenv"
)

func Init() {
	_ = godotenv.Load()
}
