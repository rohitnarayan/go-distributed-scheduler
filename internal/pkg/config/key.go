package config

import (
	"time"

	"github.com/spf13/viper"
)

func GetDurationInMillisecond(key string) time.Duration {
	return time.Millisecond * time.Duration(viper.GetInt(key))
}
