-- create "todos" table
CREATE TABLE `todos` (`id` bigint NOT NULL AUTO_INCREMENT, `text` longtext NOT NULL, `status` enum('IN_PROGRESS','COMPLETED') NOT NULL DEFAULT 'IN_PROGRESS', `created_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
