
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table person (
    id int auto_increment,

    first_name varchar(255) NOT NULL,
    last_name  varchar(255) NOT NULL,
    username   varchar(255) NOT NULL,
    password   varchar(255) NOT NULL,

    primary key (id)
);

create table student (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    person_id int NOT NULL,

    foreign key (person_id) REFERENCES person(id),
    primary key (id)
);

create table teacher (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    person_id int NOT NULL,
    course_id int,              -- this maaaay be null...

    foreign key (person_id) REFERENCES person(id),
    foreign key (course_id) REFERENCES course(id), -- does this go here or on the course?
    primary key (id)
);

create table course (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    name varchar(255) NOT NULL,
    student_id int NOT NULL,
    
    foreign key (student_id) REFERENCES student(id),
    primary key (id)      
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
DROP TABLE teacher;
DROP TABLE student;
DROP TABLE project;
DROP TABLE course;
DROP TABLE person;