package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseDSN   string `mapstructure:"DB_DSN"`
	BaseURL       string `mapstructure:"BASE_URL"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
	RedisAddr     string `mapstructure:"REDIS_ADDR"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
