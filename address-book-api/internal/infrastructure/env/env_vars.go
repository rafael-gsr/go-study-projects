// Package env contains the environment variables configurations
package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var PORT string

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

func fillVariables() {
	PORT = os.Getenv("PORT")
}
