package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	UserRegistrationEnabled bool
	Port                    int
	JwtSecret               string
	DatabaseConnectUrl      string
	LinkKeyLength           int
}

func New() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		UserRegistrationEnabled: getEnvAsBool("USER_REGISTRATION_ENABLED", true),
		Port:                    getEnvAsInt("PORT", 9000),
		JwtSecret:               getEnvOrPanic("JWT_SECRET"),
		DatabaseConnectUrl:      getEnvOrPanic("POSTGRES_CONNECT_URL"),
		LinkKeyLength:           getEnvAsInt("LINK_KEY_LENGTH", 6),
	}
}

func getEnvOrPanic(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		panic(errors.New(fmt.Sprintf("Env %s must be specified", key)))
	}

	return value
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
