package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config func to get env value from key
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
