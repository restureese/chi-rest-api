package configs

import (
	"github.com/spf13/viper"
)

type ConfigurationEnvironment struct {
	DbUri       string `mapstructure:"DB_URI"`
	JwtSecret   string `mapstructure:"JWT_SECRET"`
	JwtDuration int    `mapstructure:"JWT_DURATION"`
}

func LoadConfig(path string) (configuration ConfigurationEnvironment, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&configuration)
	return
}

var Config, _ = LoadConfig(".")
