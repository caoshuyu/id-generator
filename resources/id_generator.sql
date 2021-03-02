/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50710
 Source Host           : localhost:3306
 Source Schema         : id_generator

 Target Server Type    : MySQL
 Target Server Version : 50710
 File Encoding         : 65001

 Date: 20/02/2021 09:52:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auto_id
-- ----------------------------
DROP TABLE IF EXISTS `auto_id`;
CREATE TABLE `auto_id` (
  `id` bigint(20) unsigned NOT NULL COMMENT '标记ID',
  `project_id` varchar(64) NOT NULL COMMENT '项目标记',
  `table_name` varchar(64) NOT NULL COMMENT '表名称',
  `column_name` varchar(64) NOT NULL COMMENT '列名称',
  `st_prefix` varchar(32) NOT NULL COMMENT '前缀',
  `n_length` int(11) NOT NULL COMMENT 'Id长度',
  `st_start` int(11) NOT NULL COMMENT '开始id',
  `st_now` int(11) NOT NULL COMMENT '现在的id',
  `n_increment` int(11) NOT NULL COMMENT '步长',
  `state` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态，0正常，2禁用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ptc` (`project_id`,`table_name`,`column_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='自增id维护';

SET FOREIGN_KEY_CHECKS = 1;
