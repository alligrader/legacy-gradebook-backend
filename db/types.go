package db

import "github.com/jmoiron/sqlx"

type (
	PersonMaker struct {
		*sqlx.DB
	}

	StudentMaker struct {
		*sqlx.DB
	}

	TeacherMaker struct {
		*sqlx.DB
	}

	CourseMaker struct {
		*sqlx.DB
	}

	ProjectMaker struct {
		*sqlx.DB
	}

	SubmissionMaker struct {
		*sqlx.DB
	}
)
