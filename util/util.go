package util

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"io/ioutil"
	"strings"
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

func GetDBConfigFromEnv() *DBConfig {

	return &DBConfig{
		Flavor:     viper.GetString("DB_FLAVOR"),
		URI:        viper.GetString("DB_URI"),
		Host:       viper.GetString("DB_HOST"),
		Port:       viper.GetString("DB_PORT"),
		Name:       viper.GetString("DB_NAME"),
		SchemaFile: viper.GetString("DB_SCHEMA"),
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

func CreateTablesIfExists(db *sqlx.DB, config *DBConfig) {
	if viper.GetBool("RUNNING_SCHEMA_BROKEN") {
		ShellOut(config)
		return
	}
	RemoveDatabase(db, config.Name)
	schema := config.Schema()
	PrepAndExec(schema, db)
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

func (config *DBConfig) Schema() string {
	name := config.SchemaFile
	result, err := ioutil.ReadFile(name)
	if err != nil {
		log.Error(err)
	}
	return string(result)
}

func PrepAndExec(query string, db *sqlx.DB) error {
	var err error
	log.Info(query)
	tx, err := db.Beginx()
	defer tx.Commit()
	defer func() {
		if r := recover(); r != nil {
			log.Error(err)
			tx.Rollback()
		}
	}()

	if err != nil {
		panic(err)
	}

	stmt, err := tx.Preparex(query)
	if err != nil {
		panic(err)
	}
	if _, err = stmt.Exec(); err != nil {
		panic(err)
	}
	return nil
}

func NewestMigration() {

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

func newGooseConf() *goose.DBConf {

	var p string = viper.GetString("GOOSE_DIR")
	var env string = strings.ToLower(viper.GetString("ENV"))
	var schema string = "db"
	g, err := goose.NewDBConf(p, env, schema)
	if err != nil {
		log.Fatal(err)
	}
	return g
}
