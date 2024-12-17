package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment      string // develop, staging, production
	HttpPort         string
	LogLevel         string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	JWTSigningKey    string
	PostgresUrl      string
}

func Load() *Config {
	c := &Config{}
	path, ok := os.LookupEnv("ENV_FILE_PATH")
	if ok && path != "" {
		if err := godotenv.Load(path); err != nil {
			log.Print("No .env file found")
		}
	}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "8080"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.JWTSigningKey = cast.ToString(getOrReturnDefault("JWT_SIGNING_KEY", "secret"))

	c.PostgresUrl = cast.ToString(getOrReturnDefault("POSTGRES_URL","postgresql://postgres:password@localhost:5432/clouddb"))
	// c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	// c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	// c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "clouddb"))
	// c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	// c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "password"))

	return c
}

func getOrReturnDefault(key string, defaultValue any) any {
	v, exists := os.LookupEnv(key)
	if exists {
		return v
	}

	return defaultValue
}
