-- # mysql -u root < init.sql
-- DROP DATABASE IF EXISTS `todoList`;
-- CREATE DATABASE `todoList` CHARACTER SET utf8mb4;

# mysql -u root < init.sql
DROP DATABASE IF EXISTS `todoList`;
CREATE DATABASE `todoList` CHARACTER SET utf8mb4;
use todoList;
create user 'dbuser' @'%' identified by 'testUser';
grant all privileges on todoList.* TO 'dbuser' @'%';
