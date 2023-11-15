DROP TABLE IF EXISTS `test`.`xorm_user`;
CREATE TABLE `test`.`xorm_user` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name` varchar(60) COMMENT '名称',
    `age` tinyint(4) UNSIGNED COMMENT '年龄',
    `create_time` bigint(20) UNSIGNED COMMENT '创建时间',
    `update_time` bigint(20) UNSIGNED COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';
