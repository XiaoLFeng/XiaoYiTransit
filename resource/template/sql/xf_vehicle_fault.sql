-- 车辆故障报告表
-- 注意：执行此SQL前，请确保xf_vehicle和xf_maintenance表已经创建
CREATE TABLE IF NOT EXISTS `xf_vehicle_fault`
(
    `vehicle_fault_uuid` CHAR(36) NOT NULL COMMENT '故障UUID',
    `vehicle_uuid`       CHAR(36) NOT NULL COMMENT '车辆UUID',
    `user_uuid`          CHAR(36) NOT NULL COMMENT '报告人UUID',
    `fault_type`         VARCHAR(50) NOT NULL COMMENT '故障类型',
    `fault_description`  TEXT        NOT NULL COMMENT '故障描述',
    `report_date`        TIMESTAMP   NOT NULL COMMENT '报告时间',
    `fault_location`     VARCHAR(255)         DEFAULT NULL COMMENT '故障发生地点',
    `severity`           TINYINT(1)  NOT NULL DEFAULT '2' COMMENT '严重程度: 1-轻微, 2-一般, 3-严重, 4-危险',
    `status`             TINYINT(1)  NOT NULL DEFAULT '1' COMMENT '状态: 0-已取消, 1-待处理, 2-处理中, 3-已解决',
    `maintenance_uuid`   CHAR(36)             DEFAULT NULL COMMENT '关联的维修记录UUID',
    `notes`              TEXT COMMENT '备注',
    `created_at`         TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`         TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`vehicle_fault_uuid`),
    KEY `idx_vehicle_uuid` (`vehicle_uuid`),
    KEY `idx_user_uuid` (`user_uuid`),
    KEY `idx_report_date` (`report_date`),
    KEY `idx_status` (`status`),
    KEY `idx_maintenance_uuid` (`maintenance_uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='车辆故障报告表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_vehicle_fault`
    ADD CONSTRAINT `fk_vehicle_fault_vehicle` FOREIGN KEY (`vehicle_uuid`) REFERENCES `xf_vehicle` (`vehicle_uuid`)
        ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_vehicle_fault_maintenance` FOREIGN KEY (`maintenance_uuid`) REFERENCES `xf_maintenance` (`maintenance_uuid`)
        ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_vehicle_fault_reporter` FOREIGN KEY (`user_uuid`) REFERENCES `xf_user` (`user_uuid`)
        ON DELETE RESTRICT ON UPDATE CASCADE;
