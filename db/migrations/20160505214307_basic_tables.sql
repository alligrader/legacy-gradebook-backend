
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table t_user (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    first_name varchar(255) NOT NULL,
    last_name  varchar(255) NOT NULL,
    username   varchar(255) UNIQUE NOT NULL,
    password   varchar(255) NOT NULL,

    primary key (id)
);

create table student (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    user_id int NOT NULL,

    foreign key (user_id) REFERENCES t_user(id),
    primary key (id)
);

create table teacher (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    user_id int NOT NULL,

    foreign key (user_id) REFERENCES t_user(id),
    primary key (id)
);

create table course (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    name varchar(255) NOT NULL,
    
    primary key (id)
);

create table course_members (
    course_id int NOT NULL,
    student_id int NOT NULL,

    foreign key (course_id) REFERENCES course(id),
    foreign key (student_id) REFERENCES student(id),
    primary key(course_id, student_id)
);

create table course_teachers (
    course_id int NOT NULL,
    teacher_id int NOT NULL,

    foreign key (course_id) REFERENCES course(id),
    foreign key (teacher_id) REFERENCES teacher(id),
    primary key(course_id, teacher_id)
   
);

create table project (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    course_id int,
    name varchar(255) NOT NULL,
    description text,

    foreign key (course_id) REFERENCES course(id),
    primary key (id)
);

create table submission (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    project_id int NOT NULL,
    student_id int NOT NULL,

    foreign key (project_id) REFERENCES project(id),
    foreign key (student_id) REFERENCES student(id),
    primary key (id)
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE submission;
DROP TABLE project;
DROP TABLE course_teachers;
DROP TABLE course_members;
DROP TABLE teacher;
DROP TABLE student;
DROP TABLE course;
DROP TABLE t_user;
