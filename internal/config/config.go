package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Url      string        `yaml:"url" env-default:"http://localhost/myFirst/hs/access/"`
	Username string        `yaml:"username"  env-default:"admin"`
	Password string        `yaml:"password" env-default:""`
	Timeout  time.Duration `yaml:"timeout" env-default:"10s"`
}

func NewConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		// Возвращаем nil и ошибку
		return nil, fmt.Errorf("CONFIG_PATH env is not set (check your .env in root or system env)")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		// Оборачиваем ошибку cleanenv своим контекстом
		return nil, fmt.Errorf("cannot read config: %w", err)
	}

	return &cfg, nil
}
