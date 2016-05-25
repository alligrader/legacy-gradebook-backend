package db

import (
	"database/sql"

	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
	"github.com/jmoiron/sqlx"
	sq "gopkg.in/Masterminds/squirrel.v1"

	_ "github.com/Sirupsen/logrus"
)

func (maker *courseMaker) Create(course *Course) error {

	var (
		err                 error
		courseQuery         = queries["create_course"]
		courseMemberQuery   = queries["create_course_members"]
		courseTeachersQuery = queries["create_course_teachers"]
		acid                util.AcidTx
		result              sql.Result
		stmt                *sqlx.Stmt
	)

	acid = func(tx *sqlx.Tx) {
		// Add the course
		stmt, err = tx.Preparex(courseQuery)
		if err != nil {
			panic(err)
		}
		if result, err = stmt.Exec(course.Name); err != nil {
			panic(err)
		}
		course.ID, err = result.LastInsertId()
		if err != nil {
			panic(err)
		}

		// Then add a new record for each of the students to course_members
		for _, student := range course.Students {
			stmt, err = tx.Preparex(courseMemberQuery)
			if err != nil {
				panic(err)
			}
			if result, err = stmt.Exec(course.ID, student.ID); err != nil {
				panic(err)
			}
		}

		// Then add a new record for each of the teachers to course_teachers
		for _, teacher := range course.Teachers {
			stmt, err = tx.Preparex(courseTeachersQuery)
			if err != nil {
				panic(err)
			}
			if result, err = stmt.Exec(course.ID, teacher.ID); err != nil {
				panic(err)
			}
		}
	}
	return util.AcidCtx(acid, maker)
}

func (maker *courseMaker) UpdateCourse(course *Course) error {

	query, _, err := sq.
		Update("course").
		Set("name", course.Name).
		Where(sq.Eq{"id": course.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = util.PrepAndExec(query, maker, course.Name, course.ID)
	if err != nil {
		return err
	}

	return nil
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (maker *courseMaker) GetByID(id int64) (*Course, error) {

	var (
		err                 error
		get_course          string  = queries["get_course"]
		get_course_members  string  = queries["get_course_members"]
		get_course_teachers string  = queries["get_course_teachers"]
		course              *Course = NewCourse()
		studentIDs          []int64 = []int64{}
		teacherIDs          []int64 = []int64{}
		stmt                *sqlx.Stmt
		acid                util.AcidTx
	)

	// Get the course's name.
	acid = func(tx *sqlx.Tx) {
		// Add the course
		stmt, err = tx.Preparex(get_course)
		PanicOnError(err)

		err = stmt.Get(course, id)
		PanicOnError(err)

		// Get the rows out of course_members
		stmt, err = tx.Preparex(get_course_members)
		PanicOnError(err)

		err = stmt.Select(&studentIDs, id)
		PanicOnError(err)

		// Get the rows out of course_teachers
		stmt, err = tx.Preparex(get_course_teachers)
		PanicOnError(err)

		stmt.Select(&teacherIDs, id)
		PanicOnError(err)

		for _, studentID := range studentIDs {
			student, err := StudentStore.GetByID(studentID)
			PanicOnError(err)
			course.Students = append(course.Students, student)
		}

		for _, teacherID := range teacherIDs {
			teacher, err := TeacherStore.GetByID(teacherID)

			PanicOnError(err)
			course.Teachers = append(course.Teachers, teacher)
		}
	}
	err = util.AcidCtx(acid, maker)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (maker *courseMaker) Destroy(course *Course) error {
	query, _, err := sq.
		Delete("course").
		Where(sq.Eq{"ID": course.ID}).
		ToSql()

	if err != nil {
		return err
	}
	_, err = util.PrepAndExec(query, maker, course.ID)
	if err != nil {
		return err
	}

	return nil
}
