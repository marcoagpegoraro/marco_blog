package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvPackages() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, using the OS variables instead")
	}
}
