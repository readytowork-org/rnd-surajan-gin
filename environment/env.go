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

func GetDatabaseEnv() []string {
	return []string{os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")}
}

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}
