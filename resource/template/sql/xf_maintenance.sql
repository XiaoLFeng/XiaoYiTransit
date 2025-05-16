-- 维修记录表
-- 注意：执行此SQL前，请确保xf_vehicle表已经创建
CREATE TABLE IF NOT EXISTS `xf_maintenance`
(
    `maintenance_uuid`     CHAR(36)       NOT NULL COMMENT '维修UUID',
    `vehicle_uuid`         CHAR(36)       NOT NULL COMMENT '车辆UUID',
    `maintenance_type`     TINYINT(1)     NOT NULL COMMENT '维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修',
    `description`          text           NOT NULL COMMENT '维修描述',
    `maintenance_date`     DATE           NOT NULL COMMENT '维修日期',
    `completion_date`      DATE                    DEFAULT NULL COMMENT '完成日期',
    `cost`                 DECIMAL(10, 2) NOT NULL DEFAULT '0.00' COMMENT '维修费用',
    `mileage`              DECIMAL(10, 2)          DEFAULT NULL COMMENT '维修时里程数',
    `maintenance_location` VARCHAR(255)            DEFAULT NULL COMMENT '维修地点',
    `maintenance_staff`    VARCHAR(50)             DEFAULT NULL COMMENT '维修人员',
    `parts_replaced`       TEXT COMMENT '更换的零部件',
    `status`               TINYINT(1)     NOT NULL DEFAULT '1' COMMENT '状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成',
    `notes`                TEXT COMMENT '备注',
    `created_by_uuid`      CHAR(36)                DEFAULT NULL COMMENT '创建人UUID',
    `created_at`           TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`           TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`           TIMESTAMP               DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`maintenance_uuid`),
    KEY `idx_vehicle_uuid` (`vehicle_uuid`),
    KEY `idx_maintenance_date` (`maintenance_date`),
    KEY `idx_maintenance_type` (`maintenance_type`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='维修记录表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_maintenance`
    ADD CONSTRAINT `fk_maintenance_vehicle` FOREIGN KEY (`vehicle_uuid`) REFERENCES `xf_vehicle` (`vehicle_uuid`)
        ON DELETE RESTRICT ON UPDATE CASCADE;
