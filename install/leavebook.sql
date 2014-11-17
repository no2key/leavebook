CREATE DATABASE`leavebook`;

USE `leavebook`;

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `nick_name` varchar(10) NOT NULL COMMENT '昵称',
  `sex` tinyint(1) unsigned NOT NULL COMMENT '性别：0女1男',
  `msg` varchar(100) NOT NULL COMMENT '留言内容',
  `email` varchar(100) NOT NULL,
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ;
