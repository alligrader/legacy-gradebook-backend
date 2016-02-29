package util

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
)

var tables = []string{
	"course", "teacher", "student", "user", "test", "test_result", "run_result", "assignment", "repo",
}

func TruncateAllTables(db *sqlx.DB) {
	err := execQueryPerTable("TRUNCATE TABLE IF EXISTS %s;", db)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTablesIfNotExists(db *sqlx.DB) {
	err := execQueryPerTable("CREATE TABLE IF NOT EXISTS %s;", db)
	if err != nil {
		log.Fatal(err)
	}
}

func DropTablesIfExists(db *sqlx.DB) {
	err := execQueryPerTable("DROP TABLE IF EXISTS %s;", db)
	if err != nil {
		log.Fatal(err)
	}
}

func SetUpDatabase(db *sqlx.DB) {
	DropTablesIfExists(db)
	CreateTablesIfNotExists(db)
}

func Clean(db *sqlx.DB) {
	DropTablesIfExists(db)
}

func execQueryPerTable(query string, db *sqlx.DB) error {
	var e, err error = nil, nil

	tx, err := db.Beginx()
	defer tx.Commit()
	defer func() {
		if r := recover(); r != nil {
			e = err
			tx.Rollback()
		}
	}()

	if err != nil {
		panic(err)
	}

	for _, table := range tables {
		query := fmt.Sprintf(query, table)
		stmt, err := tx.Preparex(query)
		if err != nil {
			panic(err)
		}
		if _, err = stmt.Exec(); err != nil {
			panic(err)
		}
	}

	return e
}
