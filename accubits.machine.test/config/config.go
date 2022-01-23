package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBType                 string `mapstructure:"DB_TYPE"`
	LogFile                string `mapstructure:"LOG_FILE"`
	LogLevel               string `mapstructure:"LOG_LEVEL"`
	HostPort               string `mapstructure:"HOST_PORT"`
}

// GetConfig - Function to get Config
func GetConfig(path string) (config *Config, err error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("dev")
	v.SetConfigType("env")
	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		fmt.Println("error in reading config file.")
		return
	}
	err = v.Unmarshal(&config)
	return
}
