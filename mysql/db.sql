-- CREATE DATABASE cloud_drive CHARACTER SET utf8 COLLATE utf8_general_ci;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for file_folder
-- ----------------------------
DROP TABLE IF EXISTS `file_folder`;
CREATE TABLE `file_folder`  (
                                `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件夹ID',
                                `file_folder_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件夹名称',
                                `parent_folder_id` int(11) NULL DEFAULT 0 COMMENT '父文件夹ID',
                                `file_store_id` int(10) NULL DEFAULT NULL COMMENT '所属文件仓库ID',
                                `time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
                                PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 195 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for file_store
-- ----------------------------
DROP TABLE IF EXISTS `file_store`;
CREATE TABLE `file_store`  (
                               `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件仓库ID',
                               `user_id` int(11) NULL DEFAULT NULL COMMENT '主人ID',
                               `current_size` int(11) NULL DEFAULT 0 COMMENT '当前容量（单位KB）',
                               `max_size` int(11) NULL DEFAULT 1048576 COMMENT '最大容量（单位KB）',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 87 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for my_file
-- ----------------------------
DROP TABLE IF EXISTS `my_file`;
CREATE TABLE `my_file`  (
                            `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件ID',
                            `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件名',
                            `file_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件哈希值',
                            `file_store_id` int(10) NULL DEFAULT NULL COMMENT '文件仓库ID',
                            `download_num` int(11) NULL DEFAULT 0 COMMENT '下载次数',
                            `upload_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '上传时间',
                            `parent_folder_id` int(11)  NOT NULL DEFAULT 0 COMMENT '父文件夹ID',
                            `size` int(11) NULL DEFAULT NULL COMMENT '文件大小',
                            `type` int(11) NULL DEFAULT NULL COMMENT '文件类型',
                            `suffix` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件后缀',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 243 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for share
-- ----------------------------
DROP TABLE IF EXISTS `share`;
CREATE TABLE `share`  (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `code` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                          `file_id` int(11) NULL DEFAULT NULL,
                          `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                          `hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                         `open_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户的openid',
                         `file_store_id` int(10) NULL DEFAULT NULL COMMENT '文件仓库ID',
                         `user_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
                         `register_time` datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
                         `image_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '头像地址',
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 98 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for simple_user
-- ----------------------------
DROP TABLE IF EXISTS `simple_user`;
CREATE TABLE `simple_user`  (
                                `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                                `user_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
                                `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
                                `nick_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
                                `ext` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像拓展名',
                                `last_login_time` DATETIME NULL DEFAULT NULL COMMENT '上次登录时间',
                                PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 98 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;