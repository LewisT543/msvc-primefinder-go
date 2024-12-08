package setup

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	redisHost, err := getEnvOrError("REDIS_HOST")
	if err != nil {
		return Config{}, err
	}

	redisPort, err := getEnvOrError("REDIS_PORT")
	if err != nil {
		return Config{}, err
	}

	serverPort, err := getEnvOrError("SERVER_PORT")
	if err != nil {
		return Config{}, err
	}

	rPort, err := strconv.ParseUint(redisPort, 10, 16)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing REDIS_PORT value (%s): %v", redisPort, err)
	}

	sPort, err := strconv.ParseUint(serverPort, 10, 16)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing SERVER_PORT value (%s): %v", serverPort, err)
	}

	return Config{
		RedisAddress: fmt.Sprintf("%s:%d", redisHost, rPort),
		ServerPort:   uint16(sPort),
	}, nil
}

func getEnvOrError(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("missing required environment variable: %s", key)
	}
	return value, nil
}
