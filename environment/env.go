package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		panic("Could not load env.")
	}
}

func GetBaseUrl() string {
	return os.Getenv("BASE_URL")
}
