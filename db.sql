create database usedeal_db;

create or replace table users
(
    id         int auto_increment
        constraint `PRIMARY`
        primary key,
    username   varchar(100) not null,
    passwd     varchar(250) null,
    role       varchar(20)  null,
    first_name varchar(200) null,
    last_name  varchar(200) null,
    created_at datetime     null,
    updated_at datetime     null
);

INSERT INTO usedeall_db.users (id, username, passwd, role, first_name, last_name, created_at, updated_at) VALUES (10, 'wahyugnc', '$2a$10$KlYoceFmyBv9Zr23z/IB7eV5IEUvJTSSDI4o6aTSuVOVv1wX0qOM6', 'admin', 'Wahyu', 'Gugus Nurcahyo', '2022-09-03 08:05:22', null);
INSERT INTO usedeall_db.users (id, username, passwd, role, first_name, last_name, created_at, updated_at) VALUES (11, 'nabil123', '$2a$10$SOfVd0n12JXuq3crJw9Jfek6qZeTTbm7BLJXr9tJQujvG0cnAybx.', 'user', 'Nabil', 'Kastara', '2022-09-03 08:06:31', null);