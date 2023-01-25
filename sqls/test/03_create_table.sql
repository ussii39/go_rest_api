-- Active: 1645275120039@@127.0.0.1@3306@todoList_test
-- USE `todoList_test`;
-- DROP TABLE IF EXISTS tasks;
CREATE TABLE `todoList_test`.`tasks`(
 `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
 `uuid` varchar(255) NOT NULL DEFAULT '',
 `title` varchar(255) NOT NULL DEFAULT '',
 `detail` varchar(255) DEFAULT '' COMMENT 'titleだけでもOK',
 `user_id` bigint(20) unsigned NOT NULL,
 `status` varchar(255) NOT NULL DEFAULT 'todo' COMMENT 'todo, in_progress. in_review, done',
 `created_at` datetime DEFAULT NULL,
 `updated_at` datetime DEFAULT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- CREATE TABLE `todoList_test`.`users` (
--                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
--                                   `uuid` varchar(255) NOT NULL DEFAULT '',
--                                   `name` varchar(255) NOT NULL DEFAULT '',
--                                   `email` varchar(255) DEFAULT '' COMMENT '',
--                                   `created_at` datetime DEFAULT NULL,
--                                   `updated_at` datetime DEFAULT NULL,
--                                   PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
