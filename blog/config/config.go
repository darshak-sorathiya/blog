package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	AppPort       string
	MysqlUser     string
	MysqlDatabase string
}

func LoadAndGetConfig() *Config {
	var appConfig *Config
	if os.Getenv("ENVIRONMENT") == "test" {
		viper.SetConfigName("test")
	} else {
		viper.SetConfigName("application")
	}

	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	viper.AutomaticEnv()

	appConfig = &Config{
		AppPort:       getStringOrPanic("APP_PORT"),
		MysqlDatabase: getStringOrPanic("MYSQL_DATABASE"),
		MysqlUser:     getStringOrPanic("MYSQL_USER"),
	}
	return appConfig
}

func getStringOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}
