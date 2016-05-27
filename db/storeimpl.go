package db

import (
	"fmt"

	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
	"github.com/jmoiron/sqlx"
	sq "gopkg.in/Masterminds/squirrel.v1"

	_ "github.com/Sirupsen/logrus"
)

func (maker *userMaker) Create(user *User) error {

	query := fmt.Sprintf(queries["create_user"], user.InsertColumns())

	result, err := util.PrepAndExec(query, maker, user.FirstName, user.LastName, user.Username, string(user.Password))
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id

	return nil
}

func (maker *userMaker) GetByID(id int64) (*User, error) {

	var (
		user  *User  = &User{}
		query string = fmt.Sprintf(queries["get_user"], user.GetColumns())
		err   error  = util.GetAndMarshal(query, maker, user, id)
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (maker *studentMaker) Create(student *Student) error {

	// TODO make UserStore.Create private.
	UserStore.Create(&student.User)

	query := queries["create_student"]

	result, err := util.PrepAndExec(query, maker, student.User.ID)
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

func (maker *teacherMaker) Create(teacher *Teacher) error {
	UserStore.Create(&teacher.User)

	query, _, err := sq.
		Insert("teacher").Columns("user_id").Values("user_id").
		ToSql()
	if err != nil {
		return err
	}

	result, err := util.PrepAndExec(query, maker, teacher.User.ID)
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
		Select("teacher.id", "t_user.first_name", "t_user.last_name", "t_user.username", "t_user.created_at", "t_user.last_updated").
		From("teacher").
		Join("t_user on teacher.user_id=t_user.id").
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
