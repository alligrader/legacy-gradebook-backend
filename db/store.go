package db

type CourseStore interface {
	Ping() bool
	CreateCourse(class *Course) error
	UpdateCourse(class *Course) error
	GetCourseByID(id int) (*Course, error)
	DestroyCourse(class *Course) error
}

type TeacherStore interface {
	Ping() bool
	CreateTeacher(teacher *Teacher) error
	UpdateTeacher(teacher *Teacher) error
	GetTeacherByID(id int) (*Teacher, error)
	DestroyTeacher(t *Teacher) error
}

type StudentStore interface {
	Ping() bool
	CreateStudent(student *Student) error
	UpdateStudent(student *Student) error
	GetStudentByID(id int) (*Student, error)
	DestroyStudent(student *Student) error
}
type UserStore interface {
	Ping() bool
	CreateUser(user *User) error
	UpdateUser(user *User) error
	GetUserByID(id int) (*User, error)
	DestroyUser(user *User) error
}
type TestStore interface {
	Ping() bool
	CreateTestStore(score *TestScore) error
	UpdateTestScore(score *TestScore) error
	GetTestScoreByID(id int) (*TestScore, error)
	DestroyUser(score *TestScore) error
}
type TestResultStore interface {
	Ping() bool
	CreateTestResult(result *TestResult) error
	UpdateTestResult(result *TestResult) error
	GetTestResultByID(id int) (*TestResult, error)
	DestroyTestResult(result *TestResult) error
}
type RunResultStore interface {
	Ping() bool
	CreateRunResult(result *RunResult) error
	UpdateRunResult(result *RunResult) error
	GetRunResultByID(id int) (*RunResult, error)
	DestroyRunResult(result *RunResult) error
}
type AssignmentStore interface {
	Ping() bool
	CreateAssignment(assig *Assignment) error
	UpdateAssignment(assig *Assignment) error
	GetAssignmentByID(id int) (*Assignment, error)
	DestroyAssignment(assig *Assignment) error
}
