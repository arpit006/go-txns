package config

import (
	"fmt"
	"github.com/arpit006/go_txns/internal/constants"
	"github.com/spf13/viper"
	"log"
)

// config Application configuration
type config struct {
	appConfig AppConfig
	redisConfig RedisConfig
}

var localConfig config

// Load loads all the configurations
func Load() error {
	// initializing Viper
	viper.SetConfigName(constants.APPLICATION)
	viper.SetConfigType(constants.YAML)
	viper.AddConfigPath("configs/")
	if err := viper.ReadInConfig(); err != nil {
		// TODO log fatalf
		return err
	}

	// read config
	localConfig = config{
		appConfig: loadAppConfig(),
	}

	// TODO: validate the config
	log.Println("****LOADING****")
	log.Printf("Configuration loaded %v", localConfig)
	return nil
}

func loadAppConfig() AppConfig {
	return AppConfig{
		APP_HOST: viperString(constants.APP_HOST, "localhost"),
		APP_PORT: viperInt(constants.APP_PORT, 8080),
		LOG_LEVEL: viperString(constants.LOG_LEVEL, "INFO"),
	}
}

func loadRedisConfig() RedisConfig {
	return RedisConfig{}
}

func GetAppConfig() AppConfig { return localConfig.appConfig }
func GetAddr() string { return fmt.Sprintf("%s:%d", GetHost(), GetPort()) }
func GetHost() string { return GetAppConfig().APP_HOST }
func GetPort() int { return GetAppConfig().APP_PORT }
func GetLogLevel() string { return GetAppConfig().LOG_LEVEL }

