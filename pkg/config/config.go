package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/xd-Abi/moxie/pkg/logging"
)

func LoadEnvVariables(log *logging.Log) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Info("No .env file found")
	}

	viper.AutomaticEnv()
}

func GetUint(name string) uint {
	return viper.GetUint(name)
}

func GetString(name string) string {
	return viper.GetString(name)
}
