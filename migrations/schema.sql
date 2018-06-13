-- Adminer 4.6.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

CREATE DATABASE `api_development` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `api_development`;

DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels` (
  `ID` varchar(36) NOT NULL,
  `Name` varchar(32) NOT NULL,
  `Created_At` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Updated_At` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `channels` (`ID`, `Name`, `Created_At`, `Updated_At`) VALUES
('b83caa48-6f0f-11e8-9bb4-b46d831f881c',	'Doggo',	'2018-06-13 13:51:48',	'2018-06-13 11:41:35');

DROP TABLE IF EXISTS `channels_posts`;
CREATE TABLE `channels_posts` (
  `post_id` varchar(36) NOT NULL,
  `channel_id` varchar(36) NOT NULL,
  PRIMARY KEY (`post_id`,`channel_id`),
  KEY `channel_id` (`channel_id`),
  CONSTRAINT `channels_posts_ibfk_1` FOREIGN KEY (`post_id`) REFERENCES `posts` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `channels_posts_ibfk_2` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `channels_posts` (`post_id`, `channel_id`) VALUES
('06be6717-6f10-11e8-9bb4-b46d831f881c',	'b83caa48-6f0f-11e8-9bb4-b46d831f881c'),
('20151ada-6f10-11e8-9bb4-b46d831f881c',	'b83caa48-6f0f-11e8-9bb4-b46d831f881c'),
('a1a2fc8e-6f0f-11e8-9bb4-b46d831f881c',	'b83caa48-6f0f-11e8-9bb4-b46d831f881c');

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `ID` varchar(36) NOT NULL,
  `Likes` int(10) unsigned NOT NULL,
  `Dislikes` int(10) unsigned NOT NULL,
  `Text` varchar(255) NOT NULL,
  `Created_At` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Updated_At` timestamp NULL DEFAULT NULL,
  `post_id` varchar(36) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `post_id` (`post_id`),
  CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`post_id`) REFERENCES `posts` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `comments` (`ID`, `Likes`, `Dislikes`, `Text`, `Created_At`, `Updated_At`, `post_id`) VALUES
('e7112efa-6f0f-11e8-9bb4-b46d831f881c',	1000,	0,	'Bork',	'2018-06-13 13:51:38',	'2018-06-13 11:41:35',	'a1a2fc8e-6f0f-11e8-9bb4-b46d831f881c');

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `ID` varchar(36) NOT NULL,
  `Url` varchar(255) NOT NULL,
  `Likes` int(10) unsigned NOT NULL,
  `Dislikes` int(10) unsigned NOT NULL,
  `Status` int(11) NOT NULL,
  `Public` tinyint(4) NOT NULL,
  `Created_At` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Updated_At` timestamp NULL DEFAULT NULL,
  `user_id` varchar(36) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `posts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `posts` (`ID`, `Url`, `Likes`, `Dislikes`, `Status`, `Public`, `Created_At`, `Updated_At`, `user_id`) VALUES
('06be6717-6f10-11e8-9bb4-b46d831f881c',	'https://imgur.com/a/SihRpvO',	500,	0,	0,	1,	'2018-06-13 13:51:24',	'2018-06-13 11:41:35',	'bf2bdba2-16fb-4a4f-b17b-c98752a6e52e'),
('20151ada-6f10-11e8-9bb4-b46d831f881c',	'https://imgur.com/a/FD9TMQT',	5000,	0,	0,	1,	'2018-06-13 13:51:28',	'2018-06-13 11:41:35',	'bf2bdba2-16fb-4a4f-b17b-c98752a6e52e'),
('a1a2fc8e-6f0f-11e8-9bb4-b46d831f881c',	'https://imgur.com/a/S5h6wsc',	1000,	0,	0,	1,	'2018-06-13 13:51:32',	'2018-06-13 11:41:35',	'bf2bdba2-16fb-4a4f-b17b-c98752a6e52e');

DROP TABLE IF EXISTS `schema_migration`;
CREATE TABLE `schema_migration` (
  `version` varchar(255) NOT NULL,
  UNIQUE KEY `version_idx` (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `ID` varchar(36) NOT NULL,
  `Pseudo` varchar(32) NOT NULL,
  `Email` varchar(32) NOT NULL,
  `Password` varchar(60) NOT NULL,
  `Created_At` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Updated_At` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `users` (`ID`, `Pseudo`, `Email`, `Password`, `Created_At`, `Updated_At`) VALUES
('bf2bdba2-16fb-4a4f-b17b-c98752a6e52e',	'Adriendien',	'adrien.chevalier2@gmail.com',	'$2a$08$EYWqufyNNerX2tC.5B.Ppugu6IpOeLVUVeHqQn5iN.xF/e1uJU6BW',	'2018-06-13 11:41:35',	'2018-06-13 11:41:35');

-- 2018-06-13 13:57:58