
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table t_user (
    id           int auto_increment PRIMARY KEY,
    created_at   timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    first_name varchar(255) NOT NULL,
    last_name  varchar(255) NOT NULL,
    username   varchar(255) UNIQUE NOT NULL,
    password   varchar(255) NOT NULL,
    status     int          NOT NULL DEFAULT 0
);

-- role contains the description and UID of each role in the database
create table role (
    id          int auto_increment  PRIMARY KEY,
    name        varchar(255)        NOT NULL,
    description text                NOT NULL
);

-- role_lines maps the user.id to the role.id
create table role_lines (
    user_id int NOT NULL,
    role_id int NOT NULL,

    FOREIGN KEY (user_id) REFERENCES t_user(id),
    FOREIGN KEY (role_id) REFERENCES role(id),
    PRIMARY KEY (user_id, role_id)
);

create table action (
    id              int auto_increment PRIMARY KEY,
    title           varchar(255)       NOT NULL,
    apply_object    tinyint            NOT NULL
);

insert into action(title, apply_object) values
   ('read',     1),
   ('write',    1),
   ('delete',   1),
   ('join',     1),
   ('activate', 1),
   ('passwd',   1),
   ('list_all', 0);

-- maps the role id to an action that it can do in a certain status.
-- says a person with role X can do Y in state Z
create table t_privileges (
    role_id     int     NOT NULL,
    action_id   int     NOT NULL,
    status      int     NOT NULL,

    PRIMARY KEY (role_id, action_id, status)
);

create table student (
    id int auto_increment PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    user_id int NOT NULL,

    foreign key (user_id) REFERENCES t_user(id)
);

create table teacher (
    id int auto_increment PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    user_id int NOT NULL,

    foreign key (user_id) REFERENCES t_user(id)
);

create table course (
    id int auto_increment PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    name varchar(255) NOT NULL
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
    id int auto_increment PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    course_id int,
    name varchar(255) NOT NULL,
    description text,

    foreign key (course_id) REFERENCES course(id)
);

create table submission (
    id int auto_increment PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    project_id int NOT NULL,
    student_id int NOT NULL,

    foreign key (project_id) REFERENCES project(id),
    foreign key (student_id) REFERENCES student(id)
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE submission;
DROP TABLE project;
DROP TABLE course_teachers;
DROP TABLE course_members;
DROP TABLE course;
DROP TABLE teacher;
DROP TABLE student;
DROP TABLE t_privileges;
DROP TABLE action;
DROP TABLE role_lines;
DROP TABLE role;
DROP TABLE t_user;
