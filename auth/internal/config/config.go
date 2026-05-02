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
	pgDbname   string
}

func ReadConfig() *Config {
	cfg := &Config{}

	viper.AutomaticEnv()
	viper.SetConfigType("env")

	viper.SetDefault("APP_PORT", ":8080")

	cfg.Addr = viper.GetString("APP_PORT")
	cfg.pgHost = viper.GetString("POSTGRES_HOST")
	cfg.pgPort = viper.GetString("POSTGRES_PORT")
	cfg.pgDbname = viper.GetString("POSTGRES_DBNAME")
	cfg.pgUser = viper.GetString("POSTGRES_USER")
	cfg.pgPassword = viper.GetString("POSTGRES_PASSWORD")
	return cfg
}

func (c *Config) PostgresConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		c.pgHost,
		c.pgPort,
		c.pgDbname,
		c.pgUser,
		c.pgPassword,
	)
}
