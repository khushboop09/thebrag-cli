package configs

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LoadData() {
	err := godotenv.Load(".data")
	if err != nil {
		fmt.Println("Error loading .data file")
	}
}
