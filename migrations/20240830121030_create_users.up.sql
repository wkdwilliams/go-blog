CREATE TABLE IF NOT EXISTS `users` (
  `id` varchar(36) NOT NULL,
  `username` varchar(191) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_username` (`username`)
);