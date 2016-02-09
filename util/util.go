package util

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// This file holds the set up and tear down functions used to establish the database tables and columns.

func ClearDatabase() {

}

func CreateTablesIfExists() {

}

func Configure() {
	viper.SetEnvPrefix("SHAMAN")
	viper.AutomaticEnv()
}

func ConfigureLogger() {
	switch viper.Get("ENV") {
	case "DEVELOPMENT":
		log.SetFormatter(&log.TextFormatter{})
	case "PRODUCTION":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
}
