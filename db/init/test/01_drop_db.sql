DROP DATABASE IF EXISTS `todoList_test` CHARACTER SET utf8mb4;

CREATE DATABASE `todoList_test` CHARACTER SET utf8mb4;
use todoList_test;
create user 'dbuser'@'%' identified by 'testUser';
grant all privileges on todoList_test.* TO 'dbuser' @'%';
