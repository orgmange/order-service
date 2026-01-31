package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	Address string
	Port    string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, using system env")
	}

	version, err := getOrErr("APP_VERSION")
	if err != nil {
		return nil, err
	}

	addr, err := getOrErr("APP_ADDR")
	if err != nil {
		return nil, err
	}

	port, err := getOrErr("APP_PORT")
	if err != nil {
		return nil, err
	}

	return &Config{
		version,
		addr,
		port,
	}, nil
}

func getOrErr(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s is empty", key)
	}
	return value, nil
}
