package config

import "github.com/spf13/viper"

type Config struct {
	DSN  string `mapstructure:"dsn"`
	URL  string `mapstructure:"url"`
	Port string `mapstructure:"port"`
}

// Return a config from app.env file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
