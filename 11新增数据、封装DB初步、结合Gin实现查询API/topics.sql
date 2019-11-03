/*
Navicat MySQL Data Transfer

Source Server         : shenyi
Source Server Version : 50721
Source Host           : localhost:3306
Source Database       : gin

Target Server Type    : MYSQL
Target Server Version : 50721
File Encoding         : 65001

Date: 2019-03-07 22:21:12
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `topics`
-- ----------------------------
DROP TABLE IF EXISTS `topics`;
CREATE TABLE `topics` (
  `topic_id` int(11) NOT NULL AUTO_INCREMENT,
  `topic_title` varchar(200) NOT NULL,
  `topic_short_title` varchar(50) DEFAULT NULL,
  `user_ip` varchar(20) NOT NULL,
  `topic_score` int(11) DEFAULT NULL,
  `topic_url` varchar(200) NOT NULL,
  `topic_date` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`topic_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of topics
-- ----------------------------
INSERT INTO `topics` VALUES ('8', 'TopicTitle', 'TopicShortTitle', '127.0.0.1', '0', 'testurl', '2019-03-07 22:01:25');
