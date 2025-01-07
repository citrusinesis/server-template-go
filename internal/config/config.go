package config

import (
	"fmt"

	"go.uber.org/fx"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Host string
	Port string
}

func (sc ServerConfig) String() string {
	return fmt.Sprintf("%s:%s", sc.Host, sc.Port)
}

type DBConfig struct {
	DSN string
}

func NewConfig() (*Config, error) {
	// Load from env/file
	return &Config{
		Server: ServerConfig{
			Port: "8080",
		},
	}, nil
}

var Module = fx.Options(
	fx.Provide(NewConfig),
)
