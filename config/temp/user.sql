/*
 Navicat Premium Dump SQL

 Source Server         : 开发
 Source Server Type    : MySQL
 Source Server Version : 50744 (5.7.44-log)
 Source Host           : 43.135.97.51:3306
 Source Schema         : chat_app

 Target Server Type    : MySQL
 Target Server Version : 50744 (5.7.44-log)
 File Encoding         : 65001

 Date: 02/10/2024 22:19:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'UUID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `p_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建账号的用户ID',
  `enable` tinyint(2) NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  `last_login_time` datetime(3) NULL DEFAULT NULL COMMENT '最后登录时间',
  `token` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'token',
  `port_num` bigint(20) NULL DEFAULT 0 COMMENT ' 端口数量 也就是能开子账号的数量',
  `end_time` datetime NULL DEFAULT NULL COMMENT '会员有效期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `邮箱`(`email`) USING BTREE,
  INDEX `idx_chat_fox_user_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_admin_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for root_user
-- ----------------------------
DROP TABLE IF EXISTS `root_user`;
CREATE TABLE `root_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `enable` tinyint(2) NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  `last_login_time` datetime(3) NULL DEFAULT NULL COMMENT '最后登录时间',
  `token` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'token',
  `vip_day` bigint(20) unsigned DEFAULT '0' COMMENT '可分配会员天数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_chat_fox_user_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_admin_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;



-- ----------------------------
-- Table structure for son_user
-- ----------------------------
DROP TABLE IF EXISTS `son_user`;
CREATE TABLE `son_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'UUID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `p_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建账号的用户ID',
  `last_login_time` datetime(3) NULL DEFAULT NULL COMMENT '最后登录时间',
  `port_num` bigint(20) NULL DEFAULT 0 COMMENT ' 端口数量 也就是能开子账号的数量',
  `token` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'token',
  `enable` tinyint(2) NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UUID`(`uuid`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for vip_log
-- ----------------------------
DROP TABLE IF EXISTS `vip_log`;
CREATE TABLE `vip_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` tinyint(2) NULL DEFAULT 1 COMMENT '1.充值VIP 2.修改端口',
  `old_val` int(11) NULL DEFAULT NULL COMMENT '修改前的总账号点数',
  `new_val` int(11) NULL DEFAULT NULL COMMENT '修改后的总账号点数',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '说明',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;


INSERT INTO `root_user` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `password`, `enable`, `last_login_time`, `token`, `vip_day`) VALUES (1, NULL, '2024-10-10 19:34:05.283', NULL, 'root', '$2a$10$H7oQi.tlvksEcOPrYP8.h.aj152Iaeukkpnp28DG7oZoHo5G6J3fq', 1, '2024-10-10 19:34:05.283', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyVVVJZCI6IiIsIlVzZXJJZCI6MSwiZXhwIjoxNzI4NTg1MjQ1LCJpYXQiOjE3Mjg1NjAwNDUsImlzcyI6IldhbmdBRHVvMSIsInN1YiI6InJvb3QifQ.EonNkfrrzNdfT64-DgxrQgMhqGFEE77txkHpuU0cu2o', 365);
INSERT INTO `root_user` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `password`, `enable`, `last_login_time`, `token`, `vip_day`) VALUES (2, NULL, '2024-10-07 17:17:58.490', NULL, 'paidaxing', '$2a$10$43XqBUqJGHIA3EmZDfDyP./Itiy06BG4i3DXRECiM1.FKqBduCZJG', 1, '2024-10-07 17:17:58.489', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyVVVJZCI6IiIsIlVzZXJJZCI6MiwiZXhwIjoxNzI4MzE3ODc4LCJpYXQiOjE3MjgyOTI2NzgsImlzcyI6IldhbmdBRHVvMSIsInN1YiI6InJvb3QifQ.oTPhvHmO7jkn51Liietr4pADaZSZvlmZWn1Qv9Ann1E', 365);

