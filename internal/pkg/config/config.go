package config

import (
	"github.com/spf13/viper"
)

var (
	App *Config
)

func Init() {
	setup()

	App = &Config{
		Server: &ServerConfig{
			Name:         viper.GetString("APP_NAME"),
			Port:         viper.GetInt("HTTP_PORT"),
			ReadTimeout:  GetDurationInMillisecond("HTTP_READ_TIMEOUT_IN_MS"),
			WriteTimeout: GetDurationInMillisecond("HTTP_WRITE_TIMEOUT_IN_MS"),
		},
	}

}

func setup() {
	viper.SetConfigName("application")

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("./../../configs")
	viper.AddConfigPath("./../../../configs")
	viper.AddConfigPath("./../../../../configs")
	viper.AddConfigPath("./../../../../../configs")
	viper.AddConfigPath("./../../../../../../configs")
	viper.AddConfigPath("./../../../../../../../configs")
	viper.AddConfigPath("./../../../../../../../../configs")

	_ = viper.ReadInConfig()
	viper.AutomaticEnv()
}
