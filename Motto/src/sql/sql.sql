/*
 Navicat Premium Dump SQL

 Source Server         : 本地mysql8
 Source Server Type    : MySQL
 Source Server Version : 80041 (8.0.41)
 Source Host           : localhost:3306
 Source Schema         : sql

 Target Server Type    : MySQL
 Target Server Version : 80041 (8.0.41)
 File Encoding         : 65001

 Date: 31/07/2025 12:52:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for flag
-- ----------------------------
DROP TABLE IF EXISTS `flag`;
CREATE TABLE `flag`  (
  `thisTableSaveFlag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of flag
-- ----------------------------
INSERT INTO `flag` VALUES ('flag{sql-flag}');

-- ----------------------------
-- Table structure for motto_infos
-- ----------------------------
DROP TABLE IF EXISTS `motto_infos`;
CREATE TABLE `motto_infos`  (
  `motto_id` bigint NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `motto` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`motto_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of motto_infos
-- ----------------------------
INSERT INTO `motto_infos` VALUES (1, 'll104567', '认识的人越多 我就越喜欢狗.');
INSERT INTO `motto_infos` VALUES (3, 'Yliken', '他说你任何为人称道的美丽 不及他第一次遇到你');
INSERT INTO `motto_infos` VALUES (6, 'Yliken', '你是我患得患失的梦，我是你可有可无的人。毕竟这穿越山河的箭，刺的都是用情致疾的人');
INSERT INTO `motto_infos` VALUES (7, 'ta0', '真正的大师永远都怀一颗学徒的心');
INSERT INTO `motto_infos` VALUES (8, 'LingMj', 'The best way to learn is to teach. Keep writing, keep sharing!');
INSERT INTO `motto_infos` VALUES (9, 'c1trus', 'the quieter you become, the more you can hear');
INSERT INTO `motto_infos` VALUES (10, 'sunset', 'Records of life and study at sunset.');
INSERT INTO `motto_infos` VALUES (11, 'RedBean', '红豆生南国 春来发几枝 愿君多采撷 此物最相思');
INSERT INTO `motto_infos` VALUES (12, 'HYH', '想念的终究会相遇吧');
INSERT INTO `motto_infos` VALUES (13, 'DingTom', 'where is my shell?');


-- ----------------------------
-- Table structure for register_infos
-- ----------------------------
DROP TABLE IF EXISTS `register_infos`;
CREATE TABLE `register_infos`  (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `nickname` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `username` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `uni_register_infos_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of register_infos
-- ----------------------------
INSERT INTO `register_infos` VALUES (1, 'admin', 'admin is no use', 'admin is no use');
INSERT INTO `register_infos` VALUES (2, 'RedBean', 'RedBean', 'cannotforgetyou');

SET FOREIGN_KEY_CHECKS = 1;
