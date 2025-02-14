CREATE DATABASE IF NOT EXISTS offline_poc_db;
USE offline_poc_db;

DROP TABLE IF EXISTS notes;

CREATE TABLE notes(
    id int auto_increment primary key,
    content varchar(240) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;