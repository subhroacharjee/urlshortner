package config

import (
	"bufio"
	"fmt"
	"os"

	"github.com/subhroacharjee/urlshortner/utils/common"
)

type ConfigImpl struct {
	Environment Env
	Port        uint
	DbConfig
}

type DbConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     uint
	SSlMode  string
	TimeZone string
}

// GetDSN implements Config.
func (c *ConfigImpl) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", c.DbConfig.Host, c.DbConfig.User, c.DbConfig.Password, c.DbConfig.Name, c.DbConfig.Port, c.DbConfig.SSlMode, c.DbConfig.TimeZone)
}

// GetPort implements Config.
func (c *ConfigImpl) GetPort() uint {
	return c.Port
}

func (c ConfigImpl) GetEnv() Env {
	return c.Environment
}

func NewConfig() (Config, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }

	port, err := common.ParseUint(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	dbConfig, err := getDbConfig()
	if err != nil {
		return nil, err
	}

	return &ConfigImpl{
		Environment: Env(os.Getenv("ENV")),
		Port:        *port,
		DbConfig:    *dbConfig,
	}, nil
}

func getDbConfig() (*DbConfig, error) {
	password, err := readSecretFile("/run/secrets/db_password")
	if err != nil {
		return nil, err
	}

	user, err := readSecretFile("/run/secrets/db_user")
	if err != nil {
		return nil, err
	}

	dbPort, err := common.ParseUint(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	return &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     *user,
		Password: *password,
		Name:     os.Getenv("DB_NAME"),
		Port:     *dbPort,
		SSlMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}, nil
}

func readSecretFile(path string) (*string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	password := scanner.Text()
	return &password, nil
}
