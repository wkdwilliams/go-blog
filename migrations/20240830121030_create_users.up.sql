CREATE TABLE IF NOT EXISTS `users` (
  `id` varchar(36) NOT NULL,
  `username` varchar(50) UNIQUE NOT NULL,
  `password` binary(60) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `update_at` datetime(3) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_username` (`username`)
);