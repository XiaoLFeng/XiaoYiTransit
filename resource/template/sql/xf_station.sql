-- 站点表
CREATE TABLE IF NOT EXISTS `xf_station`
(
    `station_uuid` CHAR(36)     NOT NULL COMMENT '站点UUID',
    `name`         VARCHAR(100) NOT NULL COMMENT '站点名称',
    `code`         VARCHAR(20)           DEFAULT NULL COMMENT '站点编码',
    `address`      VARCHAR(255)          DEFAULT NULL COMMENT '站点地址',
    `longitude`    DECIMAL(10, 6)        DEFAULT NULL COMMENT '经度',
    `latitude`     DECIMAL(10, 6)        DEFAULT NULL COMMENT '纬度',
    `status`       TINYINT(1)   NOT NULL DEFAULT '1' COMMENT '状态: 0-停用, 1-启用',
    `created_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`   TIMESTAMP             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`station_uuid`),
    UNIQUE KEY `idx_name` (`name`),
    KEY `idx_code` (`code`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='站点表';