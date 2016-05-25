package db

import (
	"database/sql"

	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
	"github.com/jmoiron/sqlx"
	sq "gopkg.in/Masterminds/squirrel.v1"

	_ "github.com/Sirupsen/logrus"
)

func (maker *personMaker) Create(person *Person) error {

	query := queries["create_person"]

	result, err := util.PrepAndExec(query, maker, person.FirstName, person.LastName, person.Username, string(person.Password))
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	person.ID = id

	return nil
}

func (maker *personMaker) GetByID(id int64) (*Person, error) {

	var (
		query  string  = queries["get_person"]
		person *Person = &Person{}
		err    error   = util.GetAndMarshal(query, maker, person, id)
	)

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (maker *studentMaker) Create(student *Student) error {

	// TODO make PersonStore.Create private.
	PersonStore.Create(&student.Person)

	query := queries["create_student"]

	result, err := util.PrepAndExec(query, maker, student.Person.ID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	student.ID = id

	return nil

}
func (maker *studentMaker) Update(student *Student) error {
	return nil

}
func (maker *studentMaker) GetByID(id int64) (*Student, error) {
	var (
		student *Student = &Student{}
		query   string   = queries["get_student"]
		err     error    = util.GetAndMarshal(query, maker, student, id)
	)

	if err != nil {
		return nil, err
	}

	return student, nil
}

func (maker *studentMaker) Destroy(student *Student) error {
	return nil
}

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

func (maker *teacherMaker) Create(teacher *Teacher) error {
	PersonStore.Create(&teacher.Person)

	query, _, err := sq.
		Insert("teacher").Columns("person_id").Values("person_id").
		ToSql()
	if err != nil {
		return err
	}

	result, err := util.PrepAndExec(query, maker, teacher.Person.ID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	teacher.ID = id

	return nil
}
func (maker *teacherMaker) Update(teacher *Teacher) error {
	return nil

}
func (maker *teacherMaker) GetByID(id int64) (*Teacher, error) {
	query, _, err := sq.
		Select("teacher.id", "person.first_name", "person.last_name", "person.username", "person.created_at", "person.last_updated").
		From("teacher").
		Join("person on teacher.person_id=person.id").
		Where(sq.Eq{"teacher.id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}
	var teacher = &Teacher{}
	err = util.GetAndMarshal(query, maker, teacher, id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (maker *teacherMaker) Destroy(t *Teacher) error {
	return nil
}

type AssignmentMaker struct {
	*sqlx.DB
}

func (maker *AssignmentMaker) CreateAssignment(assig *Assignment) error {
	query, _, err := sq.
		Insert("assignment").Columns("student_id", "teacher_id").
		Values(assig.StudentID, assig.TeacherID).
		ToSql()
	if err != nil {
		return err
	}

	result, err := util.PrepAndExec(query, maker, assig.StudentID, assig.TeacherID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	assig.ID = id

	return nil
}

func (maker *AssignmentMaker) UpdateAssignment(assig *Assignment) error {
	return nil

}

func (maker *AssignmentMaker) GetAssignmentByID(id int) (*Assignment, error) {
	return nil, nil
}

func (maker *AssignmentMaker) DestroyAssignment(assig *Assignment) error {
	return nil
}
