package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type DbConfig struct {
	Host string `env:"DB_HOST, default=localhost"`
	Port string `env:"DB_PORT, default=5432"`
	User string `env:"OXYGEN_DB_USER, required"`
	Pass string `env:"OXYGEN_DB_PASS, required"`
	Name string `env:"OXYGEN_DB_NAME, required"`
}

type InitConfig struct {
	User  string `env:"OXYGEN_INIT_USER_LOGIN, default=admin"`
	Pass  string `env:"OXYGEN_INIT_USER_PASS, default=admin"`
	Name  string `env:"OXYGEN_INIT_USER_NAME, default=admin"`
	Email string `env:"OXYGEN_INIT_USER_EMAIL, default=admin@localhost"`
}

type Config struct {
	Db              DbConfig
	Init            InitConfig
	ApplicationName string `env:"APPLICATION_NAME, default=OxygenBlog"`
	Host            string `env:"HOST, default=localhost"`
	Port            string `env:"PORT, default=:8080"`
}

func ReadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process(context.Background(), &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config, nil
}
