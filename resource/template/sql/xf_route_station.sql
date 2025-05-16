-- 线路站点关联表
-- 注意：执行此SQL前，请确保xf_route和xf_station表已经创建
CREATE TABLE IF NOT EXISTS `xf_route_station`
(
    `route_station_uuid`  CHAR(36)  NOT NULL COMMENT '线路站点UUID',
    `route_uuid`          CHAR(36)  NOT NULL COMMENT '线路UUID',
    `station_uuid`        CHAR(36)  NOT NULL COMMENT '站点UUID',
    `sequence`            INT(11)   NOT NULL COMMENT '站点顺序',
    `distance_from_start` DECIMAL(10, 2)     DEFAULT '0.00' COMMENT '距起点距离(km)',
    `estimated_time`      INT(11)            DEFAULT NULL COMMENT '预计到达时间(分钟)',
    `created_at`          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`route_station_uuid`),
    UNIQUE KEY `idx_route_station_seq` (`route_uuid`, `station_uuid`, `sequence`),
    KEY `idx_station_uuid` (`station_uuid`),
    KEY `idx_route_uuid` (`route_uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='线路站点关联表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_route_station`
    ADD CONSTRAINT `fk_route_station_route` FOREIGN KEY (`route_uuid`) REFERENCES `xf_route` (`route_uuid`)
        ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_route_station_station` FOREIGN KEY (`station_uuid`) REFERENCES `xf_station` (`station_uuid`)
        ON DELETE CASCADE ON UPDATE CASCADE;
