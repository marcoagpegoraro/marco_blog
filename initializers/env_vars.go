package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvPackages() {
	if os.Getenv("ENV") != "PROD" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(".env file not found, using the OS variables instead")
		}
	}
}
