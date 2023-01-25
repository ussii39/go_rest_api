-- Active: 1645275120039@@127.0.0.1@3306@todoList
CREATE TABLE `todoList_test`.`tasks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL DEFAULT '',
  `title` varchar(255) NOT NULL DEFAULT '',
  `detail` varchar(255) DEFAULT '' COMMENT 'titleだけでもOK',
  `status` varchar(255) NOT NULL DEFAULT 'todo' COMMENT 'todo, in_progress. in_review, done',
  `user_id` bigint(20) unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

creaCREATE TABLE `todoList_test`.`households` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `cost` varchar(255) NOT NULL DEFAULT '',
  `costName` varchar(255) NOT NULL DEFAULT '',
  `isSolidCost` BOOLEAN NOT NULL DEFAULT 0,
  `user_id` bigint(20) unsigned NOT NULL,
  `resistered_at`datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `todoList_test`.`users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;