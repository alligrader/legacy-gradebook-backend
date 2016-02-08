package db

type Assignment struct{}
type Class struct{}
type School struct{}
type Teacher struct{}
type Student struct{}

type Database interface {
	Ping() bool
	CreateAssignment(assig *Assignment) error
	CreateClass(class *Class) error
	CreateSchool(school *School) error
	CreateStudent(student *Student) error
	CreateTeacher(teacher *Teacher) error

	UpdateAssignment(assig *Assignment) error
	UpdateClass(class *Class) error
	UpdateSchool(school *School) error
	UpdateStudent(student *Student) error
	UpdateTeacher(teacher *Teacher) error

	ReadAssignment(assig *Assignment) error
	ReadClass(class *Class) error
	ReadSchool(school *School) error
	ReadStudent(student *Student) error
	ReadTeacher(teacher *Teacher) error

	DestroyAssignment(assig *Assignment) error
	DestroyClass(class *Class) error
	DestroySchool(school *School) error
	DestroyStudent(student *Student) error
	DestroyTeacher(t *Teacher) error
}
