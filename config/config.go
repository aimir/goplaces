package config

import "github.com/spf13/viper"

func InitConfig() {
	viper.SetDefault("database", "places.db")
	viper.SetDefault("api_port", 8080)
	viper.SetDefault("api_allow_origin", "*")

	viper.SetConfigFile("config")
	viper.AddConfigPath("/etc/goplaces/")
	viper.AddConfigPath("/$HOME/.goplaces")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
}
