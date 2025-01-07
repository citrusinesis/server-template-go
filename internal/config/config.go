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
	Bind        BindConfig
	MaxBodySize string
}

type BindConfig struct {
	Host string
	Port string
}

func (bc BindConfig) String() string {
	return fmt.Sprintf("%s:%s", bc.Host, bc.Port)
}

type DBConfig struct {
	DSN string
}

func NewConfig() (*Config, error) {
	// Load from env/file
	return &Config{
		Server: ServerConfig{
			Bind: BindConfig{
				Port: "8080",
			},
			MaxBodySize: "1M",
		},
	}, nil
}

var Module = fx.Module("config",
	fx.Provide(NewConfig),
)
