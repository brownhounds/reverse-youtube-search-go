package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	YOUTUBE_API_KEY = "YOUTUBE_API_KEY"
	PORT            = "PORT"
)

func Init() {
	godotenv.Load() //nolint
	ValidateEnvVariables([]string{
		YOUTUBE_API_KEY,
		PORT,
	})
}

func ValidateEnvVariables(envVars []string) {
	for _, value := range envVars {
		_, defined := os.LookupEnv(value)
		if !defined {
			panic(envErrorMessage(value))
		}
	}
}

func envErrorMessage(value string) string {
	return fmt.Sprintf("ENV Variable is not defined: %s", value)
}
