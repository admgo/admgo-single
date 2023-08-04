package config

import (
	"github.com/spf13/viper"
)

func GlobalViper() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config-default")
	config.SetConfigType("yaml")
	config.AddConfigPath("config")
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("config")
	if err != nil {
		panic(err)
	}
	err = config.MergeInConfig()
	if err != nil {
		panic(err)
	}
	return config
}
