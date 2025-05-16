-- 调度表
-- 注意：执行此SQL前，请确保xf_route、xf_vehicle和xf_driver表已经创建
CREATE TABLE IF NOT EXISTS `xf_schedule`
(
    `schedule_uuid`     CHAR(36)   NOT NULL COMMENT '调度UUID',
    `route_uuid`        CHAR(36)   NOT NULL COMMENT '线路UUID',
    `vehicle_uuid`      CHAR(36)   NOT NULL COMMENT '车辆UUID',
    `driver_uuid`       CHAR(36)   NOT NULL COMMENT '司机UUID',
    `schedule_date`     DATE       NOT NULL COMMENT '调度日期',
    `start_time`        TIME       NOT NULL COMMENT '开始时间',
    `end_time`          TIME       NOT NULL COMMENT '结束时间',
    `status`            TINYINT(1) NOT NULL DEFAULT '1' COMMENT '状态: 0-取消, 1-待执行, 2-执行中, 3-已完成',
    `actual_start_time` TIME                DEFAULT NULL COMMENT '实际开始时间',
    `actual_end_time`   TIME                DEFAULT NULL COMMENT '实际结束时间',
    `mileage`           DECIMAL(10, 2)      DEFAULT '0.00' COMMENT '行驶里程(km)',
    `fuel_consumption`  DECIMAL(10, 2)      DEFAULT '0.00' COMMENT '油耗(L)',
    `passenger_count`   INT(11)             DEFAULT '0' COMMENT '载客人数',
    `notes`             TEXT COMMENT '备注',
    `created_by_uuid`   CHAR(36)            DEFAULT NULL COMMENT '创建人UUID',
    `created_at`        TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        TIMESTAMP           DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`schedule_uuid`),
    KEY `idx_route_uuid` (`route_uuid`),
    KEY `idx_vehicle_uuid` (`vehicle_uuid`),
    KEY `idx_driver_uuid` (`driver_uuid`),
    KEY `idx_schedule_date` (`schedule_date`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='调度表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_schedule`
    ADD CONSTRAINT `fk_schedule_route` FOREIGN KEY (`route_uuid`) REFERENCES `xf_route` (`route_uuid`)
        ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_schedule_vehicle` FOREIGN KEY (`vehicle_uuid`) REFERENCES `xf_vehicle` (`vehicle_uuid`)
        ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_schedule_driver` FOREIGN KEY (`driver_uuid`) REFERENCES `xf_driver` (`driver_uuid`)
        ON DELETE RESTRICT ON UPDATE CASCADE;
