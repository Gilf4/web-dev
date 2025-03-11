package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type DBConfig struct {
	USERNAME string `yaml:"USERNAME"`
	PASSWORD string `yaml:"PASSWORD"`
	PORT     string `yaml:"PORT"`
	DbName   string `yaml:"DB_NAME"`
	HOST     string `yaml:"HOST"`
}

type ServerConfig struct {
	Addr         string `yaml:"addr"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	IdleTimeout  int    `yaml:"idle_timeout"`
	Environment  string `yaml:"environment"`
}

type Config struct {
	DB     *DBConfig     `yaml:"db"`
	Server *ServerConfig `yaml:"server"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("CONFIG_PATH does not exist: %s", configPath))
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
