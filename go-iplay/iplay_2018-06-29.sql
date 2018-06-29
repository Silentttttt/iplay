# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.22)
# Database: iplay
# Generation Time: 2018-06-28 20:05:36 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table tb_choice_opt
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_choice_opt`;

CREATE TABLE `tb_choice_opt` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(512) NOT NULL DEFAULT '',
  `odds` double NOT NULL DEFAULT '0',
  `percent` double NOT NULL DEFAULT '0',
  `count` bigint(20) NOT NULL DEFAULT '0',
  `totoal` bigint(20) NOT NULL DEFAULT '0',
  `quizzes_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tb_choice_opt` WRITE;
/*!40000 ALTER TABLE `tb_choice_opt` DISABLE KEYS */;

INSERT INTO `tb_choice_opt` (`id`, `name`, `odds`, `percent`, `count`, `totoal`, `quizzes_id`)
VALUES
	(1,'乌拉圭胜',2.8,0,0,0,1),
	(2,'平',3,0,0,0,1),
	(3,'葡萄牙胜',3,0,0,0,1),
	(4,'法国胜',2.45,0,0,0,2),
	(5,'平',3,0,0,0,2),
	(6,'阿根廷胜',3.5,0,0,0,2),
	(7,'巴西胜',1.53,0,0,0,3),
	(8,'平',4.33,0,0,0,3),
	(9,'墨西哥胜',7,0,0,0,3),
	(10,'比利时胜',1.44,0,0,0,4),
	(11,'平',4,0,0,0,4),
	(12,'日本胜',8,0,0,0,4),
	(13,'西班牙胜',1.61,0,0,0,5),
	(14,'平',3.9,0,0,0,5),
	(15,'俄罗斯胜',6.5,0,0,0,5),
	(16,'克罗地亚胜',1.9,0,0,0,6),
	(17,'平',3.3,0,0,0,6),
	(18,'丹麦胜',5.2,0,0,0,6),
	(19,'瑞典胜',3.1,0,0,0,7),
	(20,'平',3,0,0,0,7),
	(21,'瑞士胜',2.7,0,0,0,7),
	(22,'哥伦比亚胜',4.2,0,0,0,8),
	(23,'平',3.3,0,0,0,8),
	(24,'英格兰胜',1.9,0,0,0,8);

/*!40000 ALTER TABLE `tb_choice_opt` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tb_game
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_game`;

CREATE TABLE `tb_game` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `play_type_id` bigint(20) NOT NULL,
  `home_team_id` bigint(20) NOT NULL,
  `visit_team_id` bigint(20) NOT NULL,
  `begin` datetime DEFAULT NULL,
  `end` datetime DEFAULT NULL,
  `description` varchar(256) DEFAULT '',
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tb_game` WRITE;
/*!40000 ALTER TABLE `tb_game` DISABLE KEYS */;

INSERT INTO `tb_game` (`id`, `play_type_id`, `home_team_id`, `visit_team_id`, `begin`, `end`, `description`, `created`)
VALUES
	(1,2,1,2,'2018-07-01 02:00:00','2018-07-01 04:00:00','世界杯淘汰赛','2018-06-26 21:00:00'),
	(2,2,3,4,'2018-06-30 22:00:00','2018-07-01 00:00:00','世界杯淘汰赛','2018-06-26 21:00:00'),
	(3,2,5,6,'2018-07-02 22:00:00','2018-07-03 00:00:00','世界杯淘汰赛','2018-06-26 21:00:00'),
	(4,2,7,8,'2018-07-03 02:00:00','2018-07-03 04:00:00','世界杯淘汰赛','2018-06-26 21:00:00'),
	(5,2,9,10,'2018-07-01 22:00:00','2018-07-02 00:00:00','世界杯淘汰赛','2018-06-26 21:00:00'),
	(6,2,11,12,'2018-07-02 02:00:00','2018-07-02 04:00:00','世界杯淘汰赛','2018-06-26 21:00:00'),
	(7,2,13,14,'2018-07-03 22:00:00','2018-07-04 00:00:00','世界杯淘汰赛','2018-06-26 21:00:00'),
	(8,2,15,16,'2018-07-04 02:00:00','2018-07-04 04:00:00','世界杯淘汰赛','2018-06-26 21:00:00');

/*!40000 ALTER TABLE `tb_game` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tb_game_ext_football
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_game_ext_football`;

CREATE TABLE `tb_game_ext_football` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` bigint(20) NOT NULL,
  `game_type` int(11) NOT NULL DEFAULT '0',
  `regular_result` int(11) NOT NULL DEFAULT '0',
  `final_result` int(11) NOT NULL DEFAULT '0',
  `home_score` int(11) NOT NULL DEFAULT '0',
  `visit_score` int(11) NOT NULL DEFAULT '0',
  `home_score_first_half` int(11) NOT NULL DEFAULT '0',
  `visit_score_first_half` int(11) NOT NULL DEFAULT '0',
  `home_score_second_half` int(11) NOT NULL DEFAULT '0',
  `visit_score_second_half` int(11) NOT NULL DEFAULT '0',
  `has_overtime` tinyint(1) NOT NULL DEFAULT '0',
  `is_overtime` tinyint(1) NOT NULL DEFAULT '0',
  `home_score_overtime` int(11) NOT NULL DEFAULT '0',
  `visit_score_overtime` int(11) NOT NULL DEFAULT '0',
  `has_penalty` tinyint(1) NOT NULL DEFAULT '0',
  `is_penalty` tinyint(1) NOT NULL DEFAULT '0',
  `home_score_penalty` int(11) NOT NULL DEFAULT '0',
  `visit_score_penalty` int(11) NOT NULL DEFAULT '0',
  `begin` datetime NOT NULL,
  `end` datetime NOT NULL,
  `description` varchar(32) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tb_play_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_play_type`;

CREATE TABLE `tb_play_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `parent` bigint(20) NOT NULL DEFAULT '0',
  `created` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tb_play_type` WRITE;
/*!40000 ALTER TABLE `tb_play_type` DISABLE KEYS */;

INSERT INTO `tb_play_type` (`id`, `name`, `parent`, `created`)
VALUES
	(1,'足球',0,NULL),
	(2,'世界杯',1,NULL);

/*!40000 ALTER TABLE `tb_play_type` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tb_player
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_player`;

CREATE TABLE `tb_player` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `play_type_id` bigint(20) NOT NULL,
  `logo` varchar(256) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tb_player` WRITE;
/*!40000 ALTER TABLE `tb_player` DISABLE KEYS */;

INSERT INTO `tb_player` (`id`, `name`, `play_type_id`, `logo`, `created`)
VALUES
	(1,'乌拉圭',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(2,'葡萄牙',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(3,'法国',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(4,'阿根廷',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(5,'巴西',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(6,'墨西哥',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(7,'比利时',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(8,'日本',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(9,'西班牙',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(10,'俄罗斯',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(11,'克罗地亚',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(12,'丹麦',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(13,'瑞典',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(14,'瑞士',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(15,'哥伦比亚',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35'),
	(16,'英格兰',2,'https://static.abcoin.pro/imgs/logo/football/536.png','2018-06-29 03:10:35');

/*!40000 ALTER TABLE `tb_player` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tb_quizzes
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_quizzes`;

CREATE TABLE `tb_quizzes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `game_id` bigint(20) NOT NULL,
  `instruction` varchar(512) NOT NULL DEFAULT '',
  `begin` datetime DEFAULT NULL,
  `end` datetime DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tb_quizzes` WRITE;
/*!40000 ALTER TABLE `tb_quizzes` DISABLE KEYS */;

INSERT INTO `tb_quizzes` (`id`, `game_id`, `instruction`, `begin`, `end`, `created`)
VALUES
	(1,1,'胜平负',NULL,'2018-07-01 02:00:00','2018-06-26 21:00:00'),
	(2,2,'胜平负',NULL,'2018-06-30 22:00:00','2018-06-26 21:00:00'),
	(3,3,'胜平负',NULL,'2018-07-02 22:00:00','2018-06-26 21:00:00'),
	(4,4,'胜平负',NULL,'2018-07-03 02:00:00','2018-06-26 21:00:00'),
	(5,5,'胜平负',NULL,'2018-07-01 22:00:00','2018-06-26 21:00:00'),
	(6,6,'胜平负',NULL,'2018-07-02 02:00:00','2018-06-26 21:00:00'),
	(7,7,'胜平负',NULL,'2018-07-03 22:00:00','2018-06-26 21:00:00'),
	(8,8,'胜平负',NULL,'2018-07-04 02:00:00','2018-06-26 21:00:00');

/*!40000 ALTER TABLE `tb_quizzes` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tb_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_user`;

CREATE TABLE `tb_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL DEFAULT '',
  `realname` varchar(32) NOT NULL DEFAULT '',
  `id_card` varchar(32) NOT NULL DEFAULT '',
  `pwd` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `mobile` varchar(16) NOT NULL DEFAULT '',
  `passphrase` varchar(255) NOT NULL DEFAULT '',
  `hash_address` varchar(256) NOT NULL DEFAULT '',
  `email` varchar(256) NOT NULL DEFAULT '',
  `avatar` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tb_user` WRITE;
/*!40000 ALTER TABLE `tb_user` DISABLE KEYS */;

INSERT INTO `tb_user` (`id`, `username`, `realname`, `id_card`, `pwd`, `status`, `mobile`, `passphrase`, `hash_address`, `email`, `avatar`)
VALUES
	(1,'','','','d41d8cd98f00b204e9800998ecf8427e',0,'','Q4ylQQz9rg','','',''),
	(2,'yangshun','','','e10adc3949ba59abbe56e057f20f883e',0,'','1QYNhyHE3z','','',''),
	(3,'leon','','','e10adc3949ba59abbe56e057f20f883e',0,'','qEqq8aejE4','','','');

/*!40000 ALTER TABLE `tb_user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tb_user_quizzes
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_user_quizzes`;

CREATE TABLE `tb_user_quizzes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `choice_opt_id` bigint(20) NOT NULL,
  `result` tinyint(1) NOT NULL DEFAULT '0',
  `money` double NOT NULL DEFAULT '0',
  `reward` double NOT NULL DEFAULT '0',
  `created` datetime NOT NULL,
  `quizzes_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
