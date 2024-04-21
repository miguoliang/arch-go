package configs

import (
	"github.com/spf13/viper"
)

func init() {

	// Set configuration file paths based on Gin mode
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../configs") // Look for configuration files in the current directory

	viper.AutomaticEnv() // Enable automatic environment variable parsing
	err := viper.ReadInConfig()
	if err != nil {
		panic("No configuration file loaded - using defaults")
	}
}
