package config

import "github.com/spf13/viper"

func viperString(config string, defaultValue ...string) string {
	if len(defaultValue) > 0 {
		viper.SetDefault(config, defaultValue[0])
	}
	return viper.GetString(config)
}

func viperInt(config string, defaultValue ...int) int {
	if len(defaultValue) > 0 {
		viper.SetDefault(config, defaultValue[0])
	}
	return viper.GetInt(config)
}
