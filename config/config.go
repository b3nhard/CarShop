package config

import (
	"os"
	"time"
)

type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}
type DbConfig struct {
	Host     string
	DbName   string
	User     string
	Password string
}

type EmailConfig struct {
	Host     string
	Username string
	Password string
	From     string
}

type Config struct {
	Server ServerConfig
	DB     DbConfig
	Email  EmailConfig
}

// Generate new configurations
func NewConfig() *Config {

	server_config := ServerConfig{
		Addr: os.Getenv("SERVER_ADDR"),
	}

	db_config := DbConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		DbName:   os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}

	email_config := EmailConfig{}

	instance := &Config{
		Server: server_config,
		DB:     db_config,
		Email:  email_config,
	}

	return instance
}
