/*
 Navicat MySQL Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : localhost:3306
 Source Schema         : college

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 24/10/2020 14:39:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for card
-- ----------------------------
DROP TABLE IF EXISTS `card`;
CREATE TABLE `card`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `totalDay` int(255) NOT NULL,
  `keepDay` int(255) NOT NULL,
  `userID` int(11) NOT NULL,
  `date` datetime(0) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of card
-- ----------------------------
INSERT INTO `card` VALUES (1, '测试', '内容测试', 45, 1, 1, '2020-04-09 08:55:02');
INSERT INTO `card` VALUES (2, '每天早起', '每天6点起床', 60, 2, 2, '2020-04-09 08:55:35');
INSERT INTO `card` VALUES (4, '吃饭', '每天按时吃饭', 12, 0, 1, '2020-04-10 17:00:09');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `shareID` int(11) NOT NULL,
  `commentType` int(11) NOT NULL,
  `userID` int(11) NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` datetime(0) NOT NULL,
  `good` int(255) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, 2, 1, 2, '太真实了。。。', '2020-04-09 10:42:41', 66);
INSERT INTO `comment` VALUES (2, 5, 2, 3, '反驳，京阿尼天下第一!', '2020-04-09 17:22:11', 4444);
INSERT INTO `comment` VALUES (3, 4, 1, 2, '测试一波', '2020-04-10 11:36:10', 0);
INSERT INTO `comment` VALUES (5, 1, 3, 1, '太厉害了', '2020-04-11 10:57:46', 0);
INSERT INTO `comment` VALUES (9, 4, 1, 5, '测试来一波', '2020-04-16 18:59:49', 1);
INSERT INTO `comment` VALUES (10, 1, 3, 5, '发送打卡感想', '2020-04-17 10:25:47', 1);
INSERT INTO `comment` VALUES (11, 2, 3, 5, '打卡感想？', '2020-04-17 10:27:51', 1);
INSERT INTO `comment` VALUES (12, 1, 1, 5, '发一条评论？', '2020-04-17 11:20:47', 1);
INSERT INTO `comment` VALUES (14, 4, 2, 6, '冰菓，男主太强了~~', '2020-04-19 17:40:08', 0);
INSERT INTO `comment` VALUES (15, 1, 3, 6, '哈哈', '2020-04-20 09:52:58', 0);
INSERT INTO `comment` VALUES (16, 5, 1, 6, '男主太强了！！！！', '2020-04-20 11:23:21', 2);
INSERT INTO `comment` VALUES (17, 11, 2, 6, '小孩子才做选择，我全都要', '2020-04-20 11:28:27', 0);
INSERT INTO `comment` VALUES (18, 5, 1, 5, '磊哥还是你磊哥', '2020-04-20 18:57:46', 0);
INSERT INTO `comment` VALUES (19, 2, 3, 5, '6点起床？？？', '2020-04-21 08:59:01', 0);
INSERT INTO `comment` VALUES (20, 4, 1, 5, '来一条评论内容', '2020-04-21 09:13:14', 0);
INSERT INTO `comment` VALUES (21, 4, 2, 5, '来一段评论内容', '2020-04-21 12:18:06', 1);
INSERT INTO `comment` VALUES (22, 4, 2, 5, '再来一段评论', '2020-04-21 12:19:09', 1);
INSERT INTO `comment` VALUES (27, 2, 1, 2, '凉凉就太真实了😂😂😂😂😂😂', '2020-04-21 17:49:27', 0);
INSERT INTO `comment` VALUES (29, 22, 2, 5, '✈️ 🚀 🖱 🐮 🌭 ⭐️ 🛵 📣 🐤 👲 😶 👳 🐶 👣 🈶 ', '2020-04-21 17:51:36', 0);

-- ----------------------------
-- Table structure for good
-- ----------------------------
DROP TABLE IF EXISTS `good`;
CREATE TABLE `good`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `userID` int(11) NOT NULL,
  `postType` int(255) NOT NULL,
  `postID` int(11) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 63 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of good
-- ----------------------------
INSERT INTO `good` VALUES (3, 1, 2, 5);
INSERT INTO `good` VALUES (4, 1, 3, 2);
INSERT INTO `good` VALUES (39, 5, 3, 7);
INSERT INTO `good` VALUES (42, 5, 3, 11);
INSERT INTO `good` VALUES (43, 5, 3, 13);
INSERT INTO `good` VALUES (46, 5, 3, 12);
INSERT INTO `good` VALUES (47, 5, 1, 1);
INSERT INTO `good` VALUES (48, 5, 2, 4);
INSERT INTO `good` VALUES (50, 6, 1, 5);
INSERT INTO `good` VALUES (51, 6, 3, 16);
INSERT INTO `good` VALUES (52, 5, 1, 5);
INSERT INTO `good` VALUES (53, 5, 3, 16);
INSERT INTO `good` VALUES (54, 5, 1, 4);
INSERT INTO `good` VALUES (55, 5, 3, 9);
INSERT INTO `good` VALUES (56, 5, 3, 10);
INSERT INTO `good` VALUES (59, 5, 3, 21);
INSERT INTO `good` VALUES (60, 5, 3, 22);
INSERT INTO `good` VALUES (61, 6, 2, 4);
INSERT INTO `good` VALUES (62, 6, 2, 11);

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `userID` int(11) NOT NULL,
  `messageType` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `postID` int(11) NOT NULL,
  `actionID` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  `date` datetime(0) NOT NULL,
  `detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `postType` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (1, 1, '4', 1, 2, 1, '2020-04-09 10:10:18', ' ', 2);
INSERT INTO `message` VALUES (2, 1, '1', 1, 3, 1, '2020-04-09 10:20:04', ' ', 2);
INSERT INTO `message` VALUES (3, 1, '2', 3, 2, 0, '2020-04-09 10:10:18', ' ', 1);
INSERT INTO `message` VALUES (4, 1, '1', 3, 3, 1, '2020-04-09 10:20:04', ' ', 1);
INSERT INTO `message` VALUES (5, 1, '1', 1, 2, 0, '2020-04-09 10:43:17', ' ', 3);
INSERT INTO `message` VALUES (6, 1, '3', 0, 0, 0, '2020-04-09 10:49:07', '你申请的xxx板块已经通过审核', 0);
INSERT INTO `message` VALUES (7, 6, '1', 5, 6, 1, '2020-04-20 11:22:18', '', 2);
INSERT INTO `message` VALUES (8, 6, '2', 5, 6, 1, '2020-04-20 11:23:21', '', 2);
INSERT INTO `message` VALUES (9, 6, '1', 16, 6, 1, '2020-04-20 11:23:44', '', 3);
INSERT INTO `message` VALUES (10, 6, '2', 11, 6, 1, '2020-04-20 11:28:27', '', 1);
INSERT INTO `message` VALUES (11, 6, '2', 18, 5, 1, '2020-04-20 18:57:46', '', 2);
INSERT INTO `message` VALUES (12, 6, '1', 5, 5, 1, '2020-04-20 18:58:27', '', 2);
INSERT INTO `message` VALUES (13, 6, '1', 16, 5, 1, '2020-04-20 18:58:31', '', 3);
INSERT INTO `message` VALUES (14, 5, '4', 4, 5, 1, '2020-04-20 18:59:09', '', 2);
INSERT INTO `message` VALUES (15, 5, '4', 9, 5, 1, '2020-04-20 18:59:11', '', 3);
INSERT INTO `message` VALUES (16, 5, '1', 10, 5, 1, '2020-04-21 09:10:55', '', 3);
INSERT INTO `message` VALUES (17, 5, '2', 20, 5, 1, '2020-04-21 09:13:14', '', 2);
INSERT INTO `message` VALUES (18, 2, '2', 21, 5, 0, '2020-04-21 12:18:06', '', 1);
INSERT INTO `message` VALUES (19, 5, '1', 21, 5, 1, '2020-04-21 12:18:13', '', 3);
INSERT INTO `message` VALUES (20, 5, '1', 21, 5, 1, '2020-04-21 12:18:16', '', 3);
INSERT INTO `message` VALUES (21, 5, '1', 21, 5, 1, '2020-04-21 12:18:49', '', 3);
INSERT INTO `message` VALUES (22, 2, '2', 22, 5, 0, '2020-04-21 12:19:09', '', 1);
INSERT INTO `message` VALUES (23, 5, '1', 22, 5, 1, '2020-04-21 12:19:11', '', 3);
INSERT INTO `message` VALUES (24, 2, '1', 4, 6, 0, '2020-04-21 15:29:49', '', 1);
INSERT INTO `message` VALUES (25, 5, '3', 11, 6, 1, '2020-04-21 15:29:56', '昊哥牛逼', 1);
INSERT INTO `message` VALUES (26, 6, '3', 1, 1, 0, '2020-04-21 16:05:53', '系统向你发送了你好啊的消息', 1);
INSERT INTO `message` VALUES (27, 5, '2', 25, 5, 0, '2020-04-21 17:45:32', '', 1);
INSERT INTO `message` VALUES (28, 6, '2', 26, 5, 0, '2020-04-21 17:47:27', '', 2);
INSERT INTO `message` VALUES (29, 6, '2', 28, 5, 0, '2020-04-21 17:49:56', '', 2);
INSERT INTO `message` VALUES (30, 5, '2', 29, 5, 0, '2020-04-21 17:51:36', '', 1);
INSERT INTO `message` VALUES (31, 1, '3', 0, 0, 0, '2020-05-05 21:16:45', '你申请的芳文社应援团板块已通过审核', 0);

-- ----------------------------
-- Table structure for plate
-- ----------------------------
DROP TABLE IF EXISTS `plate`;
CREATE TABLE `plate`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `imgUrl` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `userID` int(11) NOT NULL,
  `status` int(255) NOT NULL,
  `view` int(255) NOT NULL,
  `date` datetime(0) NOT NULL,
  `good` int(255) NOT NULL,
  `plateType` int(11) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of plate
-- ----------------------------
INSERT INTO `plate` VALUES (1, '芳文社应援团', 'https://img.xiaoyou66.com/images/2020/04/20/TrRQ.jpg', '芳文社天下第一！！！！', 1, 1, 999, '2020-04-08 09:24:08', 667, 0);
INSERT INTO `plate` VALUES (4, '印象最深刻的番剧', 'https://img.xiaoyou66.com/images/2020/03/08/2JyO.jpg', '你最喜欢看什么番剧呢？为什么----一段非常长的描述信息。。。。。', 2, 1, 8010, '2020-04-09 09:38:50', 80, 1);
INSERT INTO `plate` VALUES (5, '你最喜欢那个动漫角色', '', '就我来说我最喜欢可爱的角色，芳文社天下第一!', 1, 1, 35, '2020-04-09 16:22:31', 0, 1);
INSERT INTO `plate` VALUES (7, '测试', '', '来测试一波', 1, 0, 0, '2020-04-10 16:03:04', 0, 0);
INSERT INTO `plate` VALUES (9, '板块的名字', '', '板块的描述', 5, 0, 0, '2020-04-17 11:15:51', 0, 0);
INSERT INTO `plate` VALUES (10, '你心目中的老婆', '', '看了这么多番剧，肯定有一个心目中的老婆吧？是谁呢？', 6, 1, 16, '2020-04-20 11:27:34', 0, 1);
INSERT INTO `plate` VALUES (17, '再来测试一下发不颜文字', '', '加个表情(ง ˙o˙)ว(・ω< )★(￣へ￣)٩(๑`н´๑)۶(›´ω`‹ )( ˘•ω•˘ )(๑˙ー˙๑)(｡･ω･｡)ﾉ♡（*/∇＼*）ԅ(¯﹃¯ԅ)(´▽｀)ノ♪∠( ᐛ 」∠)＿|･ω･｀)ψ(｀∇´)ψ_(:3」∠❀)_(¦3[▓▓]', 5, 1, 2, '2020-04-21 14:50:56', 0, 1);
INSERT INTO `plate` VALUES (18, '测试使用手机发送表情包', '', '添加表情包✈️ ', 5, 1, 4, '2020-04-21 15:30:35', 0, 1);
INSERT INTO `plate` VALUES (20, '你最喜欢什么番剧', '', '什么类型的都可以说', 6, 1, 0, '2020-04-21 16:42:58', 0, 1);
INSERT INTO `plate` VALUES (21, '手机发布表情', '', '输入表情? ? ? ✈️ ? ? ? ? ? ', 5, 1, 6, '2020-04-21 17:44:23', 0, 1);
INSERT INTO `plate` VALUES (22, '手机发布表情', '', '输入表情? ✈️ ? ? ? ⭐️ ? ? ? ? ? ? ? ? ? ? ? ', 5, 1, 1, '2020-04-21 17:51:02', 0, 1);
INSERT INTO `plate` VALUES (24, '标题', '图片', '内容', 0, 1, 0, '2020-05-05 15:54:00', 0, 0);

-- ----------------------------
-- Table structure for share
-- ----------------------------
DROP TABLE IF EXISTS `share`;
CREATE TABLE `share`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `userID` int(11) NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `imgUrl` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `view` int(255) NOT NULL,
  `date` datetime(0) NOT NULL,
  `good` int(255) NOT NULL,
  `editTime` datetime(0) NOT NULL,
  `plateID` int(11) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of share
-- ----------------------------
INSERT INTO `share` VALUES (4, '分享测试+3', 5, '分享测试的内容+3', 'https://college.xiaoyou66.com/static/images/2020/4/175945.png&&https://college.xiaoyou66.com/static/images/2020/4/182321.png', 72, '2020-04-16 17:59:45', 1, '2020-04-17 10:40:07', 1);
INSERT INTO `share` VALUES (5, '冰菓观后感', 6, '冰果作为京都蓄势一年之后的诚意之作，画工配乐都可称为上乘。分镜堪称完美，画面的构图在我看来，每一帧不光富含美学方面的考虑，而且还用最合适的视角配合了故事或者感情的需要。此外，米泽靠着一部算计也多少积聚了一点人气，冰果作为其出道作并连载至今，也多少理应受到一些瞩目。但很可惜的，冰果目前还是一副不温不火的状态，甚至还比不上京都自认为的商业失败作日常。\n\n细究这种冷淡背后的原因，现在好动画很多确实无可辩驳，但更重要的是京都自己的定位的失策。冰果靠着“校园推理”这个定位，可能确实在一开始吸引了一些目光，但初看冰果，校园元素无功无过，推理元素又过于小打小闹，难免会让人失望乃至拂袖而去。但事实上，冰果的内', 'https://college.xiaoyou66.com/static/images/2020/4/112053.jpg', 34, '2020-04-20 11:20:54', 2, '2020-04-20 11:20:54', 2);
INSERT INTO `share` VALUES (6, '分享经验', 5, '经验内容', '', 1, '2020-04-21 16:59:26', 0, '2020-04-21 16:59:26', 2);
INSERT INTO `share` VALUES (8, '分享经验+2', 5, '内容', '', 0, '2020-04-21 17:02:14', 0, '2020-04-21 17:02:14', 2);

-- ----------------------------
-- Table structure for user_card
-- ----------------------------
DROP TABLE IF EXISTS `user_card`;
CREATE TABLE `user_card`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `userID` int(11) NOT NULL,
  `cardID` int(11) NOT NULL,
  `keepDay` int(11) NOT NULL,
  `lastTime` datetime(0) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_card
-- ----------------------------
INSERT INTO `user_card` VALUES (1, 1, 2, 3, '2020-04-11 10:54:43');
INSERT INTO `user_card` VALUES (2, 1, 1, 6, '2020-04-11 10:58:00');
INSERT INTO `user_card` VALUES (5, 2, 1, 0, '2020-04-11 11:03:25');
INSERT INTO `user_card` VALUES (10, 5, 2, 3, '2020-04-21 08:59:01');
INSERT INTO `user_card` VALUES (12, 6, 1, 1, '2020-04-20 09:52:58');
INSERT INTO `user_card` VALUES (14, 6, 8, 1, '2020-04-20 10:08:59');
INSERT INTO `user_card` VALUES (15, 6, 9, 1, '2020-04-20 10:09:56');
INSERT INTO `user_card` VALUES (24, 6, 2, 0, '2020-04-21 16:19:30');
INSERT INTO `user_card` VALUES (25, 5, 4, 0, '2020-04-21 16:21:47');
INSERT INTO `user_card` VALUES (27, 5, 1, 0, '2020-04-21 16:39:23');
INSERT INTO `user_card` VALUES (28, 2, 1, 0, '2020-04-20 18:09:15');

-- ----------------------------
-- Table structure for user_collect
-- ----------------------------
DROP TABLE IF EXISTS `user_collect`;
CREATE TABLE `user_collect`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `collectType` int(11) NOT NULL,
  `shareID` int(11) NOT NULL,
  `userID` int(11) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 61 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_collect
-- ----------------------------
INSERT INTO `user_collect` VALUES (1, 3, 4, 1);
INSERT INTO `user_collect` VALUES (2, 2, 2, 1);
INSERT INTO `user_collect` VALUES (5, 2, 3, 1);
INSERT INTO `user_collect` VALUES (6, 1, 1, 1);
INSERT INTO `user_collect` VALUES (20, 2, 1, 5);
INSERT INTO `user_collect` VALUES (30, 3, 4, 6);
INSERT INTO `user_collect` VALUES (31, 3, 4, 5);
INSERT INTO `user_collect` VALUES (49, 1, 1, 5);
INSERT INTO `user_collect` VALUES (51, 2, 1, 6);
INSERT INTO `user_collect` VALUES (54, 1, 2, 6);
INSERT INTO `user_collect` VALUES (56, 2, 5, 5);
INSERT INTO `user_collect` VALUES (59, 1, 2, 5);
INSERT INTO `user_collect` VALUES (60, 1, 1, 7);
INSERT INTO `user_collect` VALUES (61, 2, 4, 7);

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `imgUrl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nickName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `registeredTime` datetime(0) NOT NULL,
  `sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `college` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  UNIQUE INDEX `openid`(`openid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES (1, 'https://img.xiaoyou66.com/images/2020/02/20/tTSY.jpg', '小游', '2020-04-07 17:00:36', '男', '小游', '计算机科学与工程学院', '666');
INSERT INTO `user_info` VALUES (2, 'https://img.xiaoyou66.com/images/2020/01/21/nNUi.png', '小游', '2020-04-07 17:33:43', '保密', '无名侠', '保密', '667');
INSERT INTO `user_info` VALUES (3, 'https://img.xiaoyou66.com/images/2020/01/21/nNUi.png', '哈哈', '2020-04-08 08:18:11', '男', '小游', '化学化工学院', '668');
INSERT INTO `user_info` VALUES (4, 'https://img.xiaoyou66.com/images/2020/01/21/nNUi.png', '哈哈', '2020-04-08 15:16:01', '保密', '无名侠', '保密', '456');
INSERT INTO `user_info` VALUES (5, 'https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eqdCLCAkToM481u8IS2MH0E2UxP4nm5veUeUpuSGFBXPb6eFDvluTZUWc69keqLnTib5jeVyfzDWSw/132', '我要取一个很长很长很长的名字。', '2020-04-14 09:32:34', '保密', '有名侠一二三四五', '计算机科学与工程学院', 'otji25FTAltOF-L1mPjyYpmuz6fc');
INSERT INTO `user_info` VALUES (7, 'https://thirdwx.qlogo.cn/mmopen/vi_32/wiaNSwWd8UBUbhLpOq9O2rC6kG4vBYBibER9lEsnLQMF0Xyt4QH71H3ICY9bEg7uReegqQPGLsRTZlWEorr8rvKA/132', '小游', '2020-10-24 14:25:10', '男', '小游', '保密', 'otji25BERUFu2EFz8HIhRZlObGZ8');

SET FOREIGN_KEY_CHECKS = 1;
