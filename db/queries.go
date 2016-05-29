package db

var (
	queries = map[string]string{
		"create_user": `
			INSERT INTO t_user (%s)
			VALUES ( ?, ?, ?, ?, ?);`,
		"get_user": `
			SELECT %s
			FROM t_user
			WHERE id=?;`,
		"create_student": `
			INSERT INTO student (user_id)
			VALUES (?);`,
		"get_student": `
			SELECT student.id, t_user.id, t_user.first_name, t_user.last_name,
			t_user.username, t_user.created_at, t_user.last_updated
			FROM student
			JOIN t_user ON student.user_id=t_user.id
			WHERE student.id=?;`,
		"create_teacher": `
			INSERT INTO teacher (user_id)
			VALUES ( ? );`,
		"get_teacher": `
			SELECT teacher.id, t_user.first_name, t_user.username, 
			t_user.created_at, t_user.last_updated
			FROM teacher
			JOIN t_user ON teacher.user_id=t_user.id
			WHERE teacher.id=?;`,
		"create_course": `
			INSERT INTO course (name)
			VALUES ( ? );`,
		"create_course_members": `
			INSERT INTO course_members (course_id, student_id)
			VALUES ( ?, ? );`,
		"create_course_teachers": `
			INSERT INTO course_teachers (course_id, teacher_id)
			VALUES ( ?, ? );`,
		"get_course": `
			SELECT id, name, created_at, last_updated
			FROM course
			WHERE id=?;`,
		"get_course_members": `
			SELECT student_id
			FROM course_members
			WHERE course_id=?;`,
		"get_course_teachers": `
			SELECT teacher_id
			FROM course_teachers
			WHERE course_id=?;`,
	}
)
