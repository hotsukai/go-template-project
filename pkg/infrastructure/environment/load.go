package environment

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
