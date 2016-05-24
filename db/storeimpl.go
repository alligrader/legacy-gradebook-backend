package db

import (
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

	PersonStore.Create(&student.Person)

	query, _, err := sq.
		Insert("student").Columns("person_id").Values("person_id").
		ToSql()
	if err != nil {
		return err
	}

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
	query, _, err := sq.
		Select("student.id", "person.first_name", "person.last_name", "person.username", "person.created_at", "person.last_updated").
		From("student").
		Join("person on student.person_id=person.id").
		Where(sq.Eq{"student.id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}
	var student = &Student{}
	err = util.GetAndMarshal(query, maker, student, id)
	if err != nil {
		return nil, err
	}

	return student, nil
}
func (maker *studentMaker) Destroy(student *Student) error {
	return nil
}

func (maker *courseMaker) Create(course *Course) error {

	query, _, err := sq.
		Insert("course").Columns("name").Values("name").
		ToSql()
	if err != nil {
		return err
	}

	_ = query

	// Make a new tx
	// Add the course
	// Then add a new record for each of the students to course_members
	// Then add a new record for each of the teachers to course_teachers
	// Commit the tx
	/*
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		course.ID = int(id)
	*/
	return nil

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

func (maker *courseMaker) GetByID(id int64) (*Course, error) {
	query, _, err := sq.
		Select("id, name, created_at, last_updated").From("course").
		Where(sq.Eq{"ID": id}).
		ToSql()

	if err != nil {
		return nil, err
	}
	var course = &Course{}
	err = util.GetAndMarshal(query, maker, course, id)
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
