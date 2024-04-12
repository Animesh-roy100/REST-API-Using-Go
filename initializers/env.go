package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found. Environment variables should be present in environment.")
	}
}
