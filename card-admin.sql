/*
 Navicat Premium Data Transfer

 Source Server         : 阿里云5M服务器
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : 47.100.224.229:3306
 Source Schema         : card-admin

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 15/05/2019 10:09:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_info
-- ----------------------------
DROP TABLE IF EXISTS `admin_info`;
CREATE TABLE `admin_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限ID',
  `username` char(15) NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `head_img` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0正常  ',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=365 DEFAULT CHARSET=utf8mb4 COMMENT='管理员信息表';

-- ----------------------------
-- Table structure for admin_login_record
-- ----------------------------
DROP TABLE IF EXISTS `admin_login_record`;
CREATE TABLE `admin_login_record` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1账号密码登录    2验证码登录',
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `ip` char(20) NOT NULL DEFAULT '' COMMENT '登录IP',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip所属地址区域',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COMMENT='用户登录记录表';

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级菜单ID',
  `level` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '菜单等级 1级 2级 3级',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(20) NOT NULL DEFAULT '' COMMENT '菜单icon',
  `routers` text COMMENT '菜单路由权限',
  `sort` tinyint(4) unsigned NOT NULL DEFAULT '99' COMMENT '菜单顺序',
  `key` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单的标识',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单路由',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否是默认，默认不可删除，修改',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '0 正常 1不显示',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=743 DEFAULT CHARSET=utf8mb4 COMMENT='管理员菜单表';

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加记录的管理员ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '权限名称',
  `menu_list` text NOT NULL COMMENT '菜单列表',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否是默认',
  `is_super` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否是超级管理员菜单',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0正常',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='管理员权限角色表';

-- ----------------------------
-- Table structure for admin_token
-- ----------------------------
DROP TABLE IF EXISTS `admin_token`;
CREATE TABLE `admin_token` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0',
  `token` char(40) NOT NULL DEFAULT '',
  `ip` char(20) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0正常  ',
  `expired_time` timestamp NULL DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for pay_order
-- ----------------------------
DROP TABLE IF EXISTS `pay_order`;
CREATE TABLE `pay_order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `order_no` char(50) NOT NULL DEFAULT '' COMMENT '订单号 2018011100001',
  `service_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '服务类型，用于不同的回调业务处理',
  `service_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '服务ID',
  `title` char(255) NOT NULL DEFAULT '' COMMENT '订单标题',
  `content` char(255) NOT NULL DEFAULT '' COMMENT '订单内容描述',
  `amount` double(12,2) NOT NULL DEFAULT '0.00',
  `return_url` varchar(500) NOT NULL DEFAULT '' COMMENT '同步回调URL',
  `notify_url` varchar(500) NOT NULL DEFAULT '' COMMENT '异步回调URL',
  `pay_type` char(20) NOT NULL DEFAULT '' COMMENT '支付方式aliweb,aliqrcode,wxjs,wxqrcode,wxapp',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0未支付  1已支付',
  `pay_time` timestamp NULL DEFAULT NULL COMMENT '支付时间',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_no` (`order_no`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付订单列表';

-- ----------------------------
-- Table structure for sms_code
-- ----------------------------
DROP TABLE IF EXISTS `sms_code`;
CREATE TABLE `sms_code` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1注册账号 2解除绑定 3绑定新账号 4忘记密码  5安全校验',
  `ip` char(20) NOT NULL DEFAULT '',
  `mobile` char(20) NOT NULL COMMENT '手机号码',
  `code` char(10) NOT NULL COMMENT '验证码',
  `expire_time` int(11) unsigned NOT NULL COMMENT '有效时间',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0未使用  1已使用',
  `send_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '短信发送状态',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `mobile+type` (`mobile`,`type`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COMMENT='短信验证码表';

-- ----------------------------
-- Table structure for sms_record
-- ----------------------------
DROP TABLE IF EXISTS `sms_record`;
CREATE TABLE `sms_record` (
  `id` int(11) NOT NULL,
  `mobile` char(20) NOT NULL DEFAULT '' COMMENT '手机号码',
  `sms_code` char(20) NOT NULL DEFAULT '' COMMENT '短信码（阿里云）',
  `content` varchar(500) NOT NULL COMMENT '短信内容',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0未发送   1已发送',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `手机号码` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='短信发送记录，任意发送短信记录都会记录在该表（不管成功失败）';

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `openid` char(32) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `created_time` timestamp NULL DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `openid` char(50) NOT NULL DEFAULT '' COMMENT '微信ID',
  `username` char(15) NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `head_img` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0正常  ',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='用户记录表';

-- ----------------------------
-- Table structure for user_login_record
-- ----------------------------
DROP TABLE IF EXISTS `user_login_record`;
CREATE TABLE `user_login_record` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1账号密码登录    2验证码登录',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `ip` char(20) NOT NULL DEFAULT '' COMMENT '登录IP',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip所属地址区域',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COMMENT='用户登录记录表';

-- ----------------------------
-- Table structure for user_token
-- ----------------------------
DROP TABLE IF EXISTS `user_token`;
CREATE TABLE `user_token` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `token` char(40) NOT NULL DEFAULT '',
  `ip` char(20) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0正常  ',
  `refresh_time` timestamp NULL DEFAULT NULL,
  `expired_time` timestamp NULL DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
