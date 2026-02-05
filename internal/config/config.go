package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	Address string
	Port    string
	DBCfg   *DBCfg
}

type DBCfg struct {
	Host                    string
	User                    string
	Password                string
	Name                    string
	Port                    string
	Sslmode                 string
	MaxIdleConns            int
	MaxOpenConns            int
	MaxConnsLifetimeSeconds int
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

	DBCfg, err := LoadDBcfg()
	if err != nil {
		return nil, err
	}
	return &Config{
		version,
		addr,
		port,
		DBCfg,
	}, nil
}

func LoadDBcfg() (*DBCfg, error) {
	host, err := getOrErr("DB_HOST")
	if err != nil {
		return nil, err
	}

	user, err := getOrErr("DB_USER")
	if err != nil {
		return nil, err
	}

	pass, err := getOrErr("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	name, err := getOrErr("DB_NAME")
	if err != nil {
		return nil, err
	}

	port, err := getOrErr("DB_PORT")
	if err != nil {
		return nil, err
	}

	sslmode, err := getOrErr("DB_SSLMODE")
	if err != nil {
		return nil, err
	}

	maxIdleConns, err := getOrErrInt("DB_MAX_IDLE_CONNS")
	if err != nil {
		return nil, err
	}

	maxOpenConns, err := getOrErrInt("DB_MAX_OPEN_CONNS")
	if err != nil {
		return nil, err
	}

	maxConnsLifetimeSeconds, err := getOrErrInt("DB_CONN_MAX_LIFETIME_SECONDS")
	if err != nil {
		return nil, err
	}

	return &DBCfg{
		Host:                    host,
		User:                    user,
		Password:                pass,
		Name:                    name,
		Port:                    port,
		Sslmode:                 sslmode,
		MaxIdleConns:            maxIdleConns,
		MaxOpenConns:            maxOpenConns,
		MaxConnsLifetimeSeconds: maxConnsLifetimeSeconds,
	}, nil
}

func getOrErr(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s is empty", key)
	}
	return value, nil
}

func getOrErrInt(key string) (int, error) {
	valRaw, err := getOrErr(key)
	if err != nil {
		return 0, err
	}

	val, err := strconv.Atoi(valRaw)
	if err != nil {
		return 0, err
	}

	return val, nil
}
