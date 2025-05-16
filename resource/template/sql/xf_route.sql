-- 线路表
CREATE TABLE IF NOT EXISTS `xf_route`
(
    `route_uuid`      CHAR(36)       NOT NULL COMMENT '线路UUID',
    `route_number`    VARCHAR(20)    NOT NULL COMMENT '线路编号',
    `name`            VARCHAR(100)   NOT NULL COMMENT '线路名称',
    `start_station`   VARCHAR(100)   NOT NULL COMMENT '起始站点',
    `end_station`     VARCHAR(100)   NOT NULL COMMENT '终点站点',
    `stops`           json                    DEFAULT NULL COMMENT '途经站点(JSON格式)',
    `distance`        DECIMAL(10, 2)          DEFAULT '0.00' COMMENT '线路长度(km)',
    `fare`            DECIMAL(10, 2) NOT NULL COMMENT '票价(元)',
    `operation_hours` VARCHAR(100)            DEFAULT NULL COMMENT '运营时间',
    `frequency`       VARCHAR(50)             DEFAULT NULL COMMENT '发车频率',
    `status`          TINYINT(1)     NOT NULL DEFAULT '1' COMMENT '状态: 0-停运, 1-运营',
    `notes`           TEXT COMMENT '备注',
    `created_at`      TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`      TIMESTAMP               DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`route_uuid`),
    UNIQUE KEY `idx_route_number` (`route_number`),
    KEY `idx_name` (`name`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='线路表';
