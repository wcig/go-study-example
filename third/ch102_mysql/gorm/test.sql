DROP TABLE IF EXISTS `test`.`gorm_user`;
CREATE TABLE `test`.`gorm_user` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name` varchar(60) COMMENT '名称',
    `age` tinyint(4) UNSIGNED COMMENT '年龄',
    `create_time` bigint(20) UNSIGNED COMMENT '创建时间',
    `update_time` bigint(20) UNSIGNED COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

DROP TABLE IF EXISTS `test`.`gorm_sharding_user_0`;
CREATE TABLE `test`.`gorm_sharding_user_0` (
    `id` char(32) NOT NULL COMMENT '主键ID',
    `name` varchar(60) COMMENT '名称',
    `age` tinyint(4) UNSIGNED COMMENT '年龄',
    `create_time` bigint(20) UNSIGNED COMMENT '创建时间',
    `update_time` bigint(20) UNSIGNED COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

DROP TABLE IF EXISTS `test`.`gorm_sharding_user_1`;
CREATE TABLE `test`.`gorm_sharding_user_1` (
    `id` char(32) NOT NULL COMMENT '主键ID',
    `name` varchar(60) COMMENT '名称',
    `age` tinyint(4) UNSIGNED COMMENT '年龄',
    `create_time` bigint(20) UNSIGNED COMMENT '创建时间',
    `update_time` bigint(20) UNSIGNED COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

DROP TABLE IF EXISTS `test`.`gorm_sharding_user_2`;
CREATE TABLE `test`.`gorm_sharding_user_2` (
    `id` char(32) NOT NULL COMMENT '主键ID',
    `name` varchar(60) COMMENT '名称',
    `age` tinyint(4) UNSIGNED COMMENT '年龄',
    `create_time` bigint(20) UNSIGNED COMMENT '创建时间',
    `update_time` bigint(20) UNSIGNED COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';
