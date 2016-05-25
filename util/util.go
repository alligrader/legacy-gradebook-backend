package util

import (
	"database/sql"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	Configure()
	ConfigureLogger()
}

type AcidTx func(*sqlx.Tx)

// AcidCtx catches any panic that occurs in the call to AcidTx. If a panic occurs, tx.Rollback will be called. If not, Tx.Commit. AcidCtx injects the Tx into the func.
func AcidCtx(body AcidTx, db Execer) error {

	// Make the new Tx.
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Commit()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error("Transaction rollback.")
			err = r.(error)
			log.Error(err)
		}
	}()

	body(tx)
	return err
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

	var stmt *sqlx.Stmt
	var acid AcidTx = func(tx *sqlx.Tx) {
		stmt, err = tx.Preparex(query)
		if err != nil {
			panic(err)
		}
		if result, err = stmt.Exec(args...); err != nil {
			panic(err)
		}
	}
	AcidCtx(acid, db)

	return result, err
}

func GetAndMarshal(query string, db Execer, destination interface{}, args ...interface{}) (err error) {
	var stmt *sqlx.Stmt

	var acid AcidTx = func(tx *sqlx.Tx) {
		stmt, err = tx.Preparex(query)
		if err != nil {
			panic(err)
		}
		if err = stmt.Get(destination, args...); err != nil {
			panic(err)
		}
	}
	AcidCtx(acid, db)

	return err
}

func WithCleanDB(f func()) {
	Up()
	defer Down()
	f()
}
