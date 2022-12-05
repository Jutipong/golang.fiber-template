package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"golang.fiber.template/model"
)

var config *model.Config

func Init() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AutomaticEnv()
	v.AddConfigPath("./pkg/config")
	v.AddConfigPath("../pkg/config")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file not found")
		} else {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}

	v.Unmarshal(&config)
}

func Server() model.Server_Typd {
	return *config.Server
}

func Mssql() model.Mssql_Type {
	return *config.Mssql
}

func Ldap() model.Ldap_Type {
	return *config.Ldap
}

func Secre() model.Secret_Type {
	return *config.Secret
}

func Service() model.Service_Type {
	return *config.Service
}
