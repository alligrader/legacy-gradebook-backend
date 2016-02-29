drop database if exists shaman_test;
create database shaman_test;

create table course (
    id int auto_increment,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    primary key (id)
);

create table teacher (
    id int auto_increment,
    course_id int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    foreign key (course_id) REFERENCES course(id),
    primary key (id)
);

create table student (
    id int auto_increment,
    course_id int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    foreign key (course_id) REFERENCES course(id),
    primary key (id)
);

create table user (
    id int auto_increment,
    github_id int,
    email varchar(255),
    student_id int,
    teacher_id int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    foreign key (student_id) REFERENCES student(id),
    foreign key (teacher_id) REFERENCES teacher(id),

    primary key (id)
);

create table test (
    id int auto_increment,
    weight int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    primary key (id)
);

create table test_result (
    id int auto_increment,
    passed int,
    error_message text,
    test_id int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    foreign key (test_id) REFERENCES test(id),
    primary key (id)
);

create table run_result (
    id int auto_increment,
    test_result_id int,
    compilation_failure text,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    foreign key (test_result_id) REFERENCES test_result(id),
    primary key (id)
);

create table assignment (
    id int auto_increment,
    student_id int,
    teacher_id int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    foreign key (student_id) REFERENCES student(id),
    foreign key (teacher_id) REFERENCES teacher(id),
    primary key (id)
);

create table repo (
    id int auto_increment,
    student_id int,
    assignment_id int,
    run_result_id int,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_updated timestamp ON UPDATE CURRENT_TIMESTAMP,

    foreign key (student_id) REFERENCES student(id),
    foreign key (assignment_id) REFERENCES assignment(id),
    # foreign key (run_result_id) REFERENCES result_result(id),
    primary key (id)
);

