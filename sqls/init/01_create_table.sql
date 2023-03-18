CREATE TABLE `todoList`.`tasks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL DEFAULT '',
  `user_id` bigint(20) unsigned NOT NULL,
  `title` varchar(255) NOT NULL DEFAULT '',
  `detail` varchar(255) DEFAULT '' COMMENT 'titleだけでもOK',
  `status` varchar(255) NOT NULL DEFAULT 'todo' COMMENT 'todo, in_progress. in_review, done',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `todoList`.`users` (
                                  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                  `uuid` varchar(255) NOT NULL DEFAULT '',
                                  `name` varchar(255) NOT NULL DEFAULT '',
                                  `email` varchar(255) DEFAULT '' COMMENT '',
                                  `created_at` datetime DEFAULT NULL,
                                  `updated_at` datetime DEFAULT NULL,
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
