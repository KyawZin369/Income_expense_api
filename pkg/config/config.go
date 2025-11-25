package config

import (
	"os"
)

type Config struct {
	Port             string
	DatabaseDSN      string
	EnablePlayground bool
	JWTSecret        string
}

func Load() Config {
	port := getenv("PORT", "8080")
	dsn := getenv("DATABASE_DSN", "user:password@tcp(127.0.0.1:3306)/appdb?charset=utf8mb4&parseTime=True&loc=Local")
	playground := getenv("ENABLE_PLAYGROUND", "true") == "true"
	jwtSecret := getenv("JWT_SECRET", "your-secret-key-change-this-in-production")

	return Config{
		Port:             port,
		DatabaseDSN:      dsn,
		EnablePlayground: playground,
		JWTSecret:        jwtSecret,
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

