package main

import (
	"os"

	"github.com/matias179/go-fiber-clone/internal/server/http"

	"github.com/joho/godotenv"
)

func main() {
	envPath := os.ExpandEnv("./dotenv/.env.${APP_ENV}")

	_, err := os.Stat("/.dockerenv")

	if err == nil || os.Getenv("ECS") != "" {
		envPath = os.ExpandEnv("/bin/dotenv/.env.${APP_ENV}")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}

	http.New()
}
