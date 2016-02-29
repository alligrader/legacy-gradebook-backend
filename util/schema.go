package util

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
)

func TruncateAllTables(db *sqlx.DB) {
	err := execQueryPerTable("TRUNCATE TABLE IF EXISTS %s;", db)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTablesIfNotExists(db *sqlx.DB) {

	for _, table := range tables {
		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", table, tableDefs[table])
		if err := PrepAndExec(query, db); err != nil {
			log.Fatal(err)
		}
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

var tables = []string{
	"course", "teacher", "student", "user", "test", "test_result", "run_result", "assignment", "repo",
}

var tableDefs = map[string]string{
	"course": `
		id int auto_increment,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,
    	primary key (id)`,

	"teacher": `
		id int auto_increment,
    	course_id int,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,
    	foreign key (course_id) REFERENCES course(id),
    	primary key (id)`,

	"student": `
		id int auto_increment,
    	course_id int,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,
    	foreign key (course_id) REFERENCES course(id),
    	primary key (id)`,

	"user": `
		id int auto_increment,
    	github_id int,
    	email varchar(255),
    	student_id int,
    	teacher_id int,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,
    	foreign key (student_id) REFERENCES student(id),
    	foreign key (teacher_id) REFERENCES teacher(id),
		primary key (id)`,

	"test": `
		id int auto_increment,
    	weight int,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,
    	primary key (id)`,

	"test_result": `
		id int auto_increment,
    	passed int,
    	error_message text,
    	test_id int,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,
    	foreign key (test_id) REFERENCES test(id),
    	primary key (id)`,

	"run_result": `
		id int auto_increment,
    	test_result_id int,
    	compilation_failure text,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,
    	foreign key (test_result_id) REFERENCES test_result(id),
    	primary key (id)`,

	"assignment": `
	    id int auto_increment,
    	student_id int,
    	teacher_id int,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    	foreign key (student_id) REFERENCES student(id),
    	foreign key (teacher_id) REFERENCES teacher(id),
    	primary key (id)`,

	"repo": `
		id int auto_increment,
    	student_id int,
    	assignment_id int,
    	run_result_id int,
    	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    	last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,
    	foreign key (student_id) REFERENCES student(id),
    	foreign key (assignment_id) REFERENCES assignment(id),
    	primary key (id)`,
}
