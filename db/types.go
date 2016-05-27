package db

import (
	"github.com/alligrader/gradebook-backend/util"
	"github.com/jmoiron/sqlx"
)

func init() {
	conn := util.GetDBConfigFromEnv().ConnectToDB()
	UserStore = &userMaker{conn}
	StudentStore = &studentMaker{conn}
	TeacherStore = &teacherMaker{conn}
	CourseStore = &courseMaker{conn}
	ProjectStore = &projectMaker{conn}
	SubmissionStore = &submissionMaker{conn}
}

var (
	UserStore       *userMaker
	StudentStore    *studentMaker
	TeacherStore    *teacherMaker
	CourseStore     *courseMaker
	ProjectStore    *projectMaker
	SubmissionStore *submissionMaker
)

type (
	userMaker struct {
		*sqlx.DB
	}

	studentMaker struct {
		*sqlx.DB
	}

	teacherMaker struct {
		*sqlx.DB
	}

	courseMaker struct {
		*sqlx.DB
	}

	projectMaker struct {
		*sqlx.DB
	}

	submissionMaker struct {
		*sqlx.DB
	}
)
