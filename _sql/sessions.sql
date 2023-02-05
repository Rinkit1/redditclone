DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `token` varchar(150) NOT NULL,
  `user` int NOT NULL,
  `expires` INT UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



