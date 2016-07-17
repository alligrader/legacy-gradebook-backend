package access

import (
	"github.com/mikespook/gorbac"
)

const (
	admin     = "admin"
	assistant = "assistant"
	billing   = "billing"
	student   = "student"
	teacher   = "teacher"

	// Privileges of the Billing role
	createBilling = "create_billing_info"
	readBilling   = "read_billing_info"
	deleteBilling = "delete_billing_info"
	updateBilling = "update_billing_info"
	inviteTeacher = "invite_teacher_to_org"
	kickTeacher   = "kick_teacher_from_org"
	readCourses   = "view_courses_in_org"

	// Privileges for the Student role
	createSubmission = "create_submission_for_membered_assignment"
	readSubmission   = "review_owned_submission"
	readAssignments  = "view_assignmented_of_membered_courses"

	// Privileges of the Teacher role
	inviteStudent         = "invite_student_to_course"
	kickStudent           = "kick_student_from_course"
	readStudents          = "view_students_in_owned_courses"
	createAssignment      = "create_assignment_in_owned_course"
	readAssignment        = "read_assignment_in_owned_course"
	updateAssignment      = "update_assignment_in_owned_course"
	deleteAssignment      = "delete_assignment_in_owned_course"
	gradeSubmission       = "grade_a_submission"
	readSubmissionGrade   = "read_submission_grade_owned_course"
	updateSubmissionGrade = "update_submission_grade_owned_course"
	deleteSubmissionGrade = "delete_submission_grade"
	createCourse          = "create_course"
	updateCourse          = "update_owned_course"
	deleteCourse          = "delete_owned_course"
)

var (
	accessControl      = gorbac.New()
	Admin         Role = gorbac.NewStdRole(admin)
	Assistant     Role = gorbac.NewStdRole(assistant)
	Billing       Role = gorbac.NewStdRole(billing)
	Student       Role = gorbac.NewStdRole(student)
	Teacher       Role = gorbac.NewStdRole(teacher)

	// Privileges of the Billing role
	CreateBilling Permission = gorbac.NewStdPermission(createBilling)
	ReadBilling   Permission = gorbac.NewStdPermission(readBilling)
	DeleteBilling Permission = gorbac.NewStdPermission(deleteBilling)
	UpdateBilling Permission = gorbac.NewStdPermission(updateBilling)
	InviteTeacher Permission = gorbac.NewStdPermission(inviteTeacher)
	KickTeacher   Permission = gorbac.NewStdPermission(kickTeacher)

	// Privileges for the Student role
	CreateSubmission    Permission = gorbac.NewStdPermission(createSubmission)
	ReadSubmission      Permission = gorbac.NewStdPermission(readSubmission)
	ReadSubmissionGrade Permission = gorbac.NewStdPermission(readSubmissionGrade)
	ReadCourses         Permission = gorbac.NewStdPermission(readCourses)
	ReadAssignments     Permission = gorbac.NewStdPermission(readAssignments)

	// Privileges of the Teacher role
	InviteStudent         Permission = gorbac.NewStdPermission(inviteStudent)
	KickStudent           Permission = gorbac.NewStdPermission(kickStudent)
	ReadStudents          Permission = gorbac.NewStdPermission(readStudents)
	CreateAssignment      Permission = gorbac.NewStdPermission(createAssignment)
	ReadAssignment        Permission = gorbac.NewStdPermission(readAssignment)
	UpdateAssignment      Permission = gorbac.NewStdPermission(updateAssignment)
	DeleteAssignment      Permission = gorbac.NewStdPermission(deleteAssignment)
	GradeSubmission       Permission = gorbac.NewStdPermission(gradeSubmission)
	UpdateSubmissionGrade Permission = gorbac.NewStdPermission(updateSubmissionGrade)
	DeleteSubmissionGrade Permission = gorbac.NewStdPermission(deleteSubmissionGrade)
	CreateCourse          Permission = gorbac.NewStdPermission(createCourse)
	UpdateCourse          Permission = gorbac.NewStdPermission(readCourses)
	DeleteCourse          Permission = gorbac.NewStdPermission(deleteCourse)
)

func init() {
	assignPrivilegesToRoles()
	addRolesToACL()
	setInheritanceRelationship()
}

func setInheritanceRelationship() {
	accessControl.SetParents(admin, []string{assistant, billing, student, teacher})
}

func assignPrivilegesToRoles() {
	assignBillingPrivileges()
	assignStudentPrivileges()
	assignTeacherPrivileges()
	assignAssistantPrivileges()
}

func addRolesToACL() {
	accessControl.Add(Admin)
	accessControl.Add(Assistant)
	accessControl.Add(Billing)
	accessControl.Add(Student)
	accessControl.Add(Teacher)
}

func assignBillingPrivileges() {
	Billing.Assign(CreateBilling)
	Billing.Assign(ReadBilling)
	Billing.Assign(DeleteBilling)
	Billing.Assign(UpdateBilling)
	Billing.Assign(InviteTeacher)
	Billing.Assign(KickTeacher)
	Billing.Assign(ReadCourses)
}
func assignStudentPrivileges() {
	Student.Assign(CreateSubmission)
	Student.Assign(ReadSubmision)
	Student.Assign(ReadSubmissionGrade)
	Student.Assign(ReadCourses)
	Student.Assign(ReadAssignments)
}

func assignTeacherPrivileges() {
	Teacher.Assign(InviteStudent)
	Teacher.Assign(KickStudent)
	Teacher.Assign(ReadStudents)
	Teacher.Assign(CreateAssignment)
	Teacher.Assign(ReadAssignment)
	Teacher.Assign(UpdateAssignment)
	Teacher.Assign(DeleteAssignment)
	Teacher.Assign(GradeSubmission)
	Teacher.Assign(ReadSubmissionGrade)
	Teacher.Assign(UpdateSubmissionGrade)
	Teacher.Assign(DeleteSubmissionGrade)
	Teacher.Assign(CreateCourse)
	Teacher.Assign(ReadCourses)
	Teacher.Assign(UpdateCourse)
	Teacher.Assign(DeleteCourse)
}
func assignAssistantPrivileges() {
	Assistant.Assign(GradeSubmission)
	Assistant.Assign(ReadSubmissionGrade)
	Assistant.Assign(DeleteSubmissionGrade)
	Assistant.Assign(ReadSubmission)
	Assistant.Assign(ReadCourses)
}

type Permission gorbac.Permission
type Role gorbac.Role

func (r Role) Assign(p Permission) {
	role := gorbac.Role(r)
	perm := gorbac.Permission(p)
	role.Assign(perm)
}
