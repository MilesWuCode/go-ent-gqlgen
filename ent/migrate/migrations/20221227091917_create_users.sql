-- create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, `password` varchar(255) NOT NULL, `created_at` timestamp NOT NULL COMMENT "建立日期", `updated_at` timestamp NOT NULL COMMENT "更新日期", PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- modify "todos" table
ALTER TABLE `todos` MODIFY COLUMN `updated_at` timestamp NOT NULL, ADD COLUMN `user_id` bigint, ADD UNIQUE INDEX `user_id` (`user_id`), ADD CONSTRAINT `todos_users_todo` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
