// Package env contains the environment variables configurations
package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	env := ""

	switch os.Getenv("ENV") {
	case "production":
		env = ".production"

	case "staging":
		env = ".staging"

	default:
		env = ".development"
	}

	envfile := fmt.Sprintf(".env%s", env)
	godotenv.Load(envfile)

	fillVariables()
}

var (
	PORT string

	DB_NAME     string
	DB_TABLE    string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
)

func fillVariables() {
	PORT = os.Getenv("PORT")

	DB_NAME = os.Getenv("DB_NAME")
	DB_TABLE = os.Getenv("DB_TABLE")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	portToInt, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		DB_PORT = 0
	} else {
		DB_PORT = portToInt
	}
}
