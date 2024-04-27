package environment

import (
	"context"
)

type ApiConfig struct {
	Port       int    `env:"PORT, default=8080"`
	EnvName    string `env:"ENV_NAME, default=development"`
	ApiVersion string `env:"VERSION, default=v1"`
}

type DatabaseConfig struct {
	Url string `env:"URL, required"`
}

type Config struct {
	ApiConfig *ApiConfig      `env:",prefix=API_"`
	DbConfig  *DatabaseConfig `env:",prefix=DB_"`
}

type Environment interface {
	GetEnvironmentFromFile(ctx context.Context, fileName string) (*Config, error)
	GetEnvironment(ctx context.Context) (*Config, error)
}
