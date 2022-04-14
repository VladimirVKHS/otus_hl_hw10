CREATE USER 'otus'@'%' IDENTIFIED BY 'otus';
CREATE DATABASE otus;
GRANT ALL PRIVILEGES ON *.* TO 'otus'@'%';

USE otus;
CREATE TABLE counters (
   user_id int NOT NULL PRIMARY KEY,
   unread_messages_count int NOT NULL DEFAULT (0)
);
