package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Logger = log.New(os.Stdout, "", log.LstdFlags)

func main() {
	err := godotenv.Load()
	if err != nil {
		Logger.Println("No .env file found, checking OS environment variables...")
		for _, envVar := range []string{"HTTP_SERVER_PORT", "LIFERAY_BASE_URL", "OAUTH2_APPLICATION_REFERENCE_CODE"} {
			if os.Getenv(envVar) == "" {
				Logger.Fatalf("Required environment variable %s not found, exiting...", envVar)
			}
		}
	}

	InitHttpHandlers()

	StartHttpServer(os.Getenv("HTTP_SERVER_PORT"))
}
