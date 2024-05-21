package configs

import "github.com/spf13/viper"

type Config struct{
	PORT string
	DB_USER string
	DB_PASSWORD string
	DB_URL string
	DB_DATABASE string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}
}