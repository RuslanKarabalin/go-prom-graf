package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Addr       string
	pgUser     string
	pgPassword string
	pgHost     string
	pgPort     string
	pgDb       string
}

func ReadConfig() *Config {
	cfg := &Config{}

	viper.AutomaticEnv()
	viper.SetConfigType("env")

	viper.SetDefault("APP_PORT", ":8080")

	cfg.Addr = viper.GetString("APP_PORT")
	cfg.pgHost = viper.GetString("POSTGRES_HOST")
	cfg.pgPort = viper.GetString("POSTGRES_PORT")
	cfg.pgDb = viper.GetString("POSTGRES_DB")
	cfg.pgUser = viper.GetString("POSTGRES_USER")
	cfg.pgPassword = viper.GetString("POSTGRES_PASSWORD")
	return cfg
}

func (c *Config) PostgresConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		c.pgHost,
		c.pgPort,
		c.pgDb,
		c.pgUser,
		c.pgPassword,
	)
}
