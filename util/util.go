package util

import (
	"database/sql"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"

	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/spf13/viper/remote"
)

func init() {
	Configure()
	ConfigureLogger()
}

type DBConfig struct {
	Flavor     string
	URI        string
	Database   string
	Host       string
	Port       string
	Name       string
	SchemaFile string
}

type Execer interface {
	Beginx() (*sqlx.Tx, error)
}

func GetDBConfigFromEnv() *DBConfig {

	return &DBConfig{
		Flavor: viper.GetString("DB_FLAVOR"),
		URI:    viper.GetString("DB_URI"),
		Host:   viper.GetString("DB_HOST"),
		Port:   viper.GetString("DB_PORT"),
		Name:   viper.GetString("DB_NAME"),
	}
}

func RemoveDatabase(db *sqlx.DB, database string) {
	query := fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, database)
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		log.Error(err)
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Error(err)
	}
	if _, err := stmt.Exec(); err != nil {
		log.Error(err)
	}
	tx.Commit()
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

func (config *DBConfig) ConnectToDB() *sqlx.DB {
	result, err := sqlx.Connect(config.Flavor, config.URI)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func PrepAndExec(query string, db Execer, args ...interface{}) (result sql.Result, err error) {
	var (
		// err    error
		// result sql.Result
		tx   *sqlx.Tx
		stmt *sqlx.Stmt
	)

	log.Info(query)
	tx, err = db.Beginx()
	defer tx.Commit()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Info("We're correctly panicking and recovering")
		}
	}()

	if err != nil {
		log.Info("We're fucking up inside the commit creation")
		panic(err)
		return result, err
	}

	stmt, err = tx.Preparex(query)
	if err != nil {
		log.Info("We're fucking up inside the statement preparation")
		panic(err)
		log.Info("We're about to return from the prep statement")
		log.Info(err)
		return result, err
	}
	if result, err = stmt.Exec(args...); err != nil {
		log.Info("We're fucking up inside the stmt executation")
		panic(err)
	}

	log.Info("We're correctly returning from the function.")
	return result, err
}

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

func WithCleanDB(f func()) {
	Up()
	defer Down()
	f()
}
