package util

import (
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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
	switch viper.GetString("ENV") {
	case "DEVELOPMENT":
		log.SetFormatter(&log.TextFormatter{})
	case "PRODUCTION":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
}

func ConnectToDB() *sqlx.DB {
	driver := viper.GetString("DB_FLAVOR")
	uri := viper.GetString("DB_URI")
	result, err := sqlx.Connect(driver, uri)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
