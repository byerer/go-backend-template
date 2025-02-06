package config

import (
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var serveConfig *GlobalConfig

func LoadConfig(configYml string) {
	serveConfig = new(GlobalConfig)
	viper.SetConfigFile(configYml)
	err := viper.ReadInConfig()
	if err != nil {
		println("Config Read failed: " + err.Error())
		os.Exit(1)
	}
	err = viper.Unmarshal(serveConfig)
	if err != nil {
		println("Config Unmarshal failed: " + err.Error())
		os.Exit(1)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		println("Config fileHandle changed: ", e.Name)
		_ = viper.ReadInConfig()
		err = viper.Unmarshal(serveConfig)
		if err != nil {
			println("New Config fileHandle Parse Failed: ", e.Name)
			return
		}
	})
	viper.WatchConfig()
}

func GetConfig() *GlobalConfig {
	return serveConfig
}
