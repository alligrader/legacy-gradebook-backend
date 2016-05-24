package db

var (
	queries = map[string]string{
		"create_person": `
			INSERT INTO person (first_name, last_name, username, password)
			VALUES ( ?, ?, ?, ?);`,
		"get_person": `
			SELECT id, first_name, last_name, username, created_at, last_updated 
			FROM person
			WHERE id=?;`,
		"create_student": `
			INSERT INTO student (person_id)
			VALUES (?);`,
		"get_student": `
			SELECT student.id, person.id, person.last_name, 
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
	}
)
