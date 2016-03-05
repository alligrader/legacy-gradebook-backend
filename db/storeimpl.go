package db

import (
	. "github.com/gradeshaman/gradebook-backend/models"
	"github.com/gradeshaman/gradebook-backend/util"
	"github.com/jmoiron/sqlx"
	sq "gopkg.in/Masterminds/squirrel.v1"

	log "github.com/Sirupsen/logrus"
)

type CourseMaker struct {
	*sqlx.DB
}

func (maker *CourseMaker) CreateCourse(course *Course) error {

	query, _, err := sq.
		Insert("course").Columns().Values().
		ToSql()
	if err != nil {
		return err
	}

	log.Info(query)

	result, err := util.PrepAndExec(query, maker)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	course.ID = int(id)

	return nil

}

func (maker *CourseMaker) UpdateCourse(class *Course) error {

	return nil
}

func (maker *CourseMaker) GetCourseByID(id int) (*Course, error) {
	query, _, err := sq.
		Select("id, created_at, last_updated").From("course").
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
func (maker *CourseMaker) DestroyCourse(class *Course) error {

	return nil
}

type TeacherMaker struct {
	*sqlx.DB
}

func (maker *TeacherMaker) CreateTeacher(teacher *Teacher) error {
	return nil

}
func (maker *TeacherMaker) UpdateTeacher(teacher *Teacher) error {
	return nil

}
func (maker *TeacherMaker) GetTeacherByID(id int) (*Teacher, error) {
	return nil, nil

}

func (maker *TeacherMaker) DestroyTeacher(t *Teacher) error {
	return nil
}

type StudentMaker struct {
	*sqlx.DB
}

func (maker *StudentMaker) CreateStudent(student *Student) error {
	return nil

}
func (maker *StudentMaker) UpdateStudent(student *Student) error {
	return nil

}
func (maker *StudentMaker) GetStudentByID(id int) (*Student, error) {
	return nil, nil

}
func (maker *StudentMaker) DestroyStudent(student *Student) error {
	return nil
}

type UserMaker struct {
	*sqlx.DB
}

func (maker *UserMaker) CreateUser(user *User) error {
	return nil

}
func (maker *UserMaker) UpdateUser(user *User) error {
	return nil

}
func (maker *UserMaker) GetUserByID(id int) (*User, error) {
	return nil, nil

}
func (maker *UserMaker) DestroyUser(user *User) error {
	return nil

}

type TestScoreMaker struct {
	*sqlx.DB
}

func (maker *TestScoreMaker) CreateTestStore(score *Test) error {
	return nil

}
func (maker *TestScoreMaker) UpdateTestScore(score *Test) error {
	return nil

}
func (maker *TestScoreMaker) GetTestScoreByID(id int) (*Test, error) {
	return nil, nil

}
func (maker *TestScoreMaker) DestroyUser(score *Test) error {
	return nil
}

type TestResultMaker struct {
	*sqlx.DB
}

func (maker *TestResultMaker) CreateTestResult(result *TestResult) error {
	return nil

}
func (maker *TestResultMaker) UpdateTestResult(result *TestResult) error {
	return nil

}
func (maker *TestResultMaker) GetTestResultByID(id int) (*TestResult, error) {
	return nil, nil

}
func (maker *TestResultMaker) DestroyTestResult(result *TestResult) error {
	return nil
}

type RunResultMaker struct {
	*sqlx.DB
}

func (maker *RunResultMaker) CreateRunResult(result *RunResult) error {
	return nil

}
func (maker *RunResultMaker) UpdateRunResult(result *RunResult) error {
	return nil

}
func (maker *RunResultMaker) GetRunResultByID(id int) (*RunResult, error) {
	return nil, nil

}
func (maker *RunResultMaker) DestroyRunResult(result *RunResult) error {
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
	log.Warn(query)

	result, err := util.PrepAndExec(query, maker, assig.StudentID, assig.TeacherID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	assig.ID = int(id)

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
