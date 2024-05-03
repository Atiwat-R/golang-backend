package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port string

	DBUser string
	DBPassword string
	DBAddress string
	DBName string
}

var Envs = initConfig()

func initConfig() Config {

	// Load .env environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "172.17.0.2"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "golang-db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

