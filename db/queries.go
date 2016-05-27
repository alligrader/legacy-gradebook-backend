package db

var (
	queries = map[string]string{
		"create_person": `
			INSERT INTO person (%s)
			VALUES ( ?, ?, ?, ?);`,
		"get_person": `
			SELECT %s
			FROM person
			WHERE id=?;`,
		"create_student": `
			INSERT INTO student (person_id)
			VALUES (?);`,
		"get_student": `
			SELECT student.id, person.id, person.first_name, person.last_name,
			person.username, person.created_at, person.last_updated
			FROM student
			JOIN person ON student.person_id=person.id
			WHERE student.id=?;`,
		"create_teacher": `
			INSERT INTO teacher (person_id)
			VALUES ( ? );`,
		"get_teacher": `
			SELECT teacher.id, person.first_name, person.username, 
			person.created_at, person.last_updated
			FROM teacher
			JOIN person ON teacher.person_id=person.id
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
