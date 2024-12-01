package util

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL    string `mapstructure:"DATABASE_URL"`
	DBUsername     string `mapstructure:"DB_USERNAME"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	Source         string `mapstructure:"SOURCE"`
	Environment    string `mapstructure:"ENVIRONMENT"`
	HttpServerAddr string `mapstructure:"HTTP_SERVER_ADDR"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))

	}

	err = viper.Unmarshal(&config)
	return
}
