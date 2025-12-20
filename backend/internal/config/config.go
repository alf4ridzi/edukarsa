package config

import "github.com/spf13/viper"

type Config struct {
	ServerPort          int    `mapstructure:"SERVER_PORT"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBName              string `mapstructure:"DB_NAME"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBPort              int    `mapstructure:"DB_PORT"`
	AccessSecret        string `mapstructure:"ACCESS_SECRET"`
	RefreshSecret       string `mapstructure:"REFRESH_SECRET"`
	AccessTokenExpired  int    `mapstructure:"ACCESS_TOKEN_EXPIRED"`
	RefreshTokenExpired int    `mapstructure:"REFRESH_TOKEN_EXPIRED"`
	ServerHost          string `mapstructure:"SERVER_HOST"`
}

// global
var AppConfig *Config

func LoadConfig() (*Config, error) {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return &config, err
}
