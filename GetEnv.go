package main

import (
	"os"

	dotenv "github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := dotenv.Load(".env")
	HandleError(err)

	return os.Getenv(key)
}
