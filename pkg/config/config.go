package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"golang.fiber.template/model"
)

func Init() (config *model.Config) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetConfigFile("config.yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file not found")
		} else {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}

	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
