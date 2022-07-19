/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : imall

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 19/07/2022 19:16:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '地址id',
  `open_id` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '微信用户id',
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '收货人姓名',
  `mobile` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '手机号',
  `province` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '省份',
  `city` char(30) DEFAULT NULL COMMENT '城市',
  `district` char(30) DEFAULT NULL COMMENT '区/县',
  `detailed_address` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '详细地址',
  `is_default` tinyint DEFAULT NULL COMMENT '1为默认，2为非默认',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1112 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of address
-- ----------------------------
BEGIN;
INSERT INTO `address` (`id`, `open_id`, `name`, `mobile`, `province`, `city`, `district`, `detailed_address`, `is_default`, `created`, `updated`) VALUES (1109, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '张先生', '13364809604', '北京市', '北京市', '朝阳区', '和平街都20号楼205', 1, '2022-06-23 14:42:46', '');
INSERT INTO `address` (`id`, `open_id`, `name`, `mobile`, `province`, `city`, `district`, `detailed_address`, `is_default`, `created`, `updated`) VALUES (1110, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '李先生', '13580609702', '山西省', '阳泉市', '郊区', '某某街都2508号5栋450', 2, '2022-06-23 14:43:50', '');
INSERT INTO `address` (`id`, `open_id`, `name`, `mobile`, `province`, `city`, `district`, `detailed_address`, `is_default`, `created`, `updated`) VALUES (1111, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', 'zchengo', '15060894035', '浙江省', '杭州市', '上城区', '某某街道708号', 2, '2022-06-23 14:44:32', '2022-06-23 14:44:42');
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '类目编号',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '类目名称',
  `parent_id` bigint DEFAULT NULL COMMENT '父级类目编号',
  `level` tinyint DEFAULT NULL COMMENT '类目级别',
  `sort` int DEFAULT NULL COMMENT '类目排序',
  `created` varchar(50) DEFAULT NULL COMMENT '创建时间',
  `updated` varchar(50) DEFAULT NULL COMMENT '更新时间',
  `sid` bigint DEFAULT NULL COMMENT '店铺编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2077 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2055, '蔬菜', 1, 1, 1, '2022-06-23 11:48:33', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2056, '水果', 1, 1, 2, '2022-06-23 11:48:43', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2057, '菜心', 2055, 2, 20, '2022-06-23 11:48:57', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2058, '上海青', 2055, 2, 23, '2022-06-23 11:55:27', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2061, '火龙果', 2056, 2, 34, '2022-06-23 12:15:23', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2062, '柠檬', 2056, 2, 33, '2022-06-23 12:38:43', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2063, '苹果', 2056, 2, 12, '2022-06-23 12:58:50', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2064, '香蕉', 2056, 2, 5, '2022-06-23 13:04:19', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2065, '香梨', 2056, 2, 23, '2022-06-23 13:18:29', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2066, '西柚', 2056, 2, 55, '2022-06-23 13:44:04', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2067, '西瓜', 2056, 2, 23, '2022-06-23 13:54:53', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2068, '甜瓜', 2056, 2, 22, '2022-06-23 14:06:43', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2069, '秋葵', 2055, 2, 30, '2022-06-23 14:10:10', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2070, '胡萝卜', 2055, 2, 60, '2022-06-23 14:13:16', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2071, '土豆', 2055, 2, 33, '2022-06-23 14:15:37', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2072, '玉米', 2055, 2, 12, '2022-06-23 14:17:39', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2073, '海鲜菇', 2055, 2, 5, '2022-06-23 14:20:47', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2074, '洋葱', 2055, 2, 23, '2022-06-23 14:22:43', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2075, '香菜', 2055, 2, 50, '2022-06-23 14:24:59', '', 100001);
INSERT INTO `category` (`id`, `name`, `parent_id`, `level`, `sort`, `created`, `updated`, `sid`) VALUES (2076, '菜椒', 2055, 2, 45, '2022-06-23 14:29:18', '', 100001);
COMMIT;

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品编号',
  `category_id` bigint DEFAULT NULL COMMENT '类目编号',
  `title` varchar(250) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品标题',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品名称',
  `price` decimal(10,2) DEFAULT NULL COMMENT '商品价格',
  `quantity` int DEFAULT NULL COMMENT '商品数量',
  `image_url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品图片',
  `remark` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品备注',
  `sales` int DEFAULT NULL COMMENT '商品销量',
  `status` tinyint DEFAULT NULL COMMENT '商品状态',
  `created` varchar(50) DEFAULT NULL COMMENT '创建时间',
  `updated` varchar(50) DEFAULT NULL COMMENT '更新时间',
  `sid` bigint DEFAULT NULL COMMENT '店铺编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of goods
-- ----------------------------
BEGIN;
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (36, 2057, '中花菜心 500g', '菜心', 5.00, 500, 'http://localhost:8000/image/IMG_0001.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 11:52:44', '2022-06-23 11:52:54', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (37, 2058, '有机上海青 300g', '上海青', 8.00, 330, 'http://localhost:8000/image/IMG_0002.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 11:56:18', '2022-06-23 11:56:31', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (39, 2061, '红心火龙果 200g', '火龙果', 13.00, 30, 'http://localhost:8000/image/IMG_0003.PNG', '图片仅供参考，请以实物为准', 0, 2, '2022-06-23 12:26:53', '2022-06-23 12:57:55', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (40, 2062, '安岳柠檬 120g', '柠檬', 3.00, 340, 'http://localhost:8000/image/IMG_0004.PNG', '图片仅供参考，请以实物为准', 0, 2, '2022-06-23 12:47:36', '2022-06-23 12:58:22', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (41, 2063, '南非冰糖小苹果 450g', '苹果', 10.00, 220, 'http://localhost:8000/image/IMG_0005.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 13:02:01', '2022-07-11 21:57:38', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (42, 2064, '香蕉 500g', '香蕉', 8.00, 450, 'http://localhost:8000/image/IMG_0006.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 13:05:19', '', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (43, 2065, '新疆香梨 400g', '香梨', 10.00, 23, 'http://localhost:8000/image/IMG_0007.JPG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 13:34:35', '', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (44, 2066, '红西柚 500g', '红西柚', 7.00, 60, 'http://localhost:8000/image/IMG_0008.JPG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 13:44:59', '', 100002);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (45, 2067, '麒麟瓜 5kg', '麒麟瓜', 32.00, 200, 'http://localhost:8000/image/IMG_0009.JPG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 13:56:55', '', 100002);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (46, 2068, '黄甜瓜 600g', '甜瓜', 9.00, 350, 'http://localhost:8000/image/IMG_0010.JPG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:08:37', '', 100002);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (47, 2069, '黄秋葵 250g', '黄秋葵', 5.00, 50, 'http://localhost:8000/image/IMG_0011.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:12:18', '', 100002);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (48, 2070, '胡萝卜 260g', '胡萝卜', 3.00, 105, 'http://localhost:8000/image/IMG_0012.JPG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:14:45', '', 100002);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (49, 2071, '土豆 200g', '土豆', 2.00, 68, 'http://localhost:8000/image/IMG_0013.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:17:05', '', 100002);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (50, 2072, '黄羊河甜玉米 250g', '甜玉米', 8.00, 200, 'http://localhost:8000/image/IMG_0014.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:19:21', '', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (51, 2073, '海鲜菇 240g', '海鲜菇', 3.00, 30, 'http://localhost:8000/image/IMG_0015.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:21:49', '', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (52, 2074, '洋葱 200g', '洋葱', 4.00, 200, 'http://localhost:8000/image/IMG_0016.JPG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:24:11', '', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (53, 2075, '香菜 90g', '香菜', 3.00, 12, 'http://localhost:8000/image/IMG_0017.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:26:52', '', 100001);
INSERT INTO `goods` (`id`, `category_id`, `title`, `name`, `price`, `quantity`, `image_url`, `remark`, `sales`, `status`, `created`, `updated`, `sid`) VALUES (54, 2076, '菜椒 300g', '菜椒', 4.00, 300, 'http://localhost:8000/image/IMG_0018.PNG', '图片仅供参考，请以实物为准', 0, 1, '2022-06-23 14:30:54', '', 100001);
COMMIT;

-- ----------------------------
-- Table structure for market
-- ----------------------------
DROP TABLE IF EXISTS `market`;
CREATE TABLE `market` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `type` tinyint DEFAULT NULL COMMENT '类型',
  `banner_image` varchar(200) DEFAULT NULL COMMENT '活动图片',
  `begin_time` varchar(50) DEFAULT NULL COMMENT '开始时间',
  `over_time` varchar(50) DEFAULT NULL COMMENT '结束时间',
  `goods_ids` varchar(250) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '关联商品',
  `status` tinyint DEFAULT NULL COMMENT '状态，1-开启，2-关闭',
  `created` varchar(50) DEFAULT NULL COMMENT '创建时间',
  `updated` varchar(50) DEFAULT NULL COMMENT '更新时间',
  `sid` bigint DEFAULT NULL COMMENT '店铺编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of market
-- ----------------------------
BEGIN;
INSERT INTO `market` (`id`, `name`, `type`, `banner_image`, `begin_time`, `over_time`, `goods_ids`, `status`, `created`, `updated`, `sid`) VALUES (11, '钻石展位1', 1, 'http://localhost:8000/image/b1.png', '2022-06-23 00:00:00', '2022-06-30 00:00:00', '42', 1, '2022-06-23 14:33:28', '2022-07-18 20:28:12', 100001);
INSERT INTO `market` (`id`, `name`, `type`, `banner_image`, `begin_time`, `over_time`, `goods_ids`, `status`, `created`, `updated`, `sid`) VALUES (12, '钻石展位2', 1, 'http://localhost:8000/image/b2.png', '2022-06-23 00:00:00', '2022-06-30 00:00:00', '40', 1, '2022-06-23 14:34:11', '2022-07-19 17:55:28', 100001);
COMMIT;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '订单编号',
  `open_id` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '微信用户编号',
  `goods_ids_count` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品项Id',
  `goods_count` int DEFAULT NULL COMMENT '商品数量',
  `total_price` decimal(10,2) DEFAULT NULL COMMENT '订单金额',
  `address_id` bigint DEFAULT NULL COMMENT '地址编号',
  `status` tinyint DEFAULT NULL COMMENT '订单状态，1-待付款，2-已取消，3-已付款，4-配送中，5-已完成',
  `created` varchar(50) DEFAULT NULL COMMENT '创建时间',
  `updated` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '更新时间',
  `sid` bigint DEFAULT NULL COMMENT '店铺编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of order
-- ----------------------------
BEGIN;
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (12, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 1, '2022-07-19 06:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (13, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2', 7, 58.00, 1109, 2, '2022-07-19 06:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (14, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 3, '2022-07-19 08:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (15, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 5, '2022-07-19 08:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (16, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3', 7, 58.00, 1109, 1, '2022-07-19 10:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (17, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 2, '2022-07-19 14:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (18, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 3, '2022-07-19 14:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (19, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 3, '2022-07-19 14:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (20, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 4, '2022-07-19 15:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (22, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 1, '2022-07-19 18:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (23, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:3,43:2,44:2', 7, 58.00, 1109, 3, '2022-07-19 18:48:40', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (39, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '40:3', 3, 9.00, 1109, 2, '2022-06-26 14:53:02', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (40, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '36:2,37:2', 4, 26.00, 1109, 4, '2022-06-26 14:53:39', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (41, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '53:2', 2, 6.00, 1109, 3, '2022-06-26 14:55:03', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (42, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:2,50:2', 4, 32.00, 1109, 4, '2022-06-27 14:55:14', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (43, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '42:2,49:4,50:3', 9, 48.00, 1109, 5, '2022-06-27 14:55:52', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (44, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '52:4,53:4', 8, 28.00, 1109, 4, '2022-06-27 14:56:02', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (45, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '36:3', 3, 15.00, 1109, 1, '2022-06-27 14:56:11', '', 100001);
INSERT INTO `order` (`id`, `open_id`, `goods_ids_count`, `goods_count`, `total_price`, `address_id`, `status`, `created`, `updated`, `sid`) VALUES (47, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '37:4', 4, 32.00, 1109, 1, '2022-07-19 17:42:03', '', 100001);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `open_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '微信用户唯一标识',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户名称',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户密码',
  `status` tinyint DEFAULT NULL COMMENT '用户状态',
  `created` varchar(50) DEFAULT NULL COMMENT '创建时间',
  `updated` varchar(50) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100004 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `open_id`, `username`, `password`, `status`, `created`, `updated`) VALUES (100001, NULL, 'admin', '12345', 1, '2022-07-20 14:36:10', NULL);
INSERT INTO `user` (`id`, `open_id`, `username`, `password`, `status`, `created`, `updated`) VALUES (100002, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', 'zchengo', '12345', 1, '2022-06-23 14:36:10', '');
INSERT INTO `user` (`id`, `open_id`, `username`, `password`, `status`, `created`, `updated`) VALUES (100003, '', '', '', 1, '2022-07-19 17:39:18', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
