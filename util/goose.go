package util

import (
	"fmt"
	"strings"

	"bitbucket.org/liamstask/goose/lib/goose"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

func newGooseConf() *goose.DBConf {

	var p string = viper.GetString("GOOSE_DIR")
	var env string = strings.ToLower(viper.GetString("ENV"))
	var schema string = "db"
	g, err := goose.NewDBConf(p, env, schema)
	if err != nil {
		panic(err)
	}
	return g
}

func Up() {

	var dirpath string = viper.GetString("GOOSE_DIR")
	cfg := newGooseConf()
	fmt.Printf("Open str: %s", cfg.Driver.OpenStr)
	version, err := goose.GetMostRecentDBVersion(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	err = goose.RunMigrations(cfg, dirpath, version)
	if err != nil {
		log.Fatal(err)
	}
}

func Down() {
	var dirpath string = viper.GetString("GOOSE_DIR")
	cfg := newGooseConf()
	version, err := goose.GetMostRecentDBVersion(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	previous, err := goose.GetPreviousDBVersion(dirpath, version)
	if err != nil {
		log.Fatal(err)
	}
	err = goose.RunMigrations(cfg, dirpath, previous)
	if err != nil {
		log.Fatal(err)
	}
}
