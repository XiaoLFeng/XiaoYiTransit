-- 车辆表
CREATE TABLE IF NOT EXISTS `xf_vehicle`
(
    `vehicle_uuid`          CHAR(36)    NOT NULL COMMENT '车辆UUID',
    `plate_number`          VARCHAR(20) NOT NULL COMMENT '车牌号',
    `model`                 VARCHAR(50) NOT NULL COMMENT '车辆型号',
    `purchase_date`         DATE        NOT NULL COMMENT '购买日期',
    `status`                TINYINT(1)  NOT NULL DEFAULT '1' COMMENT '状态: 1-运营, 2-维修, 3-停运, 4-报废',
    `seats`                 INT(11)     NOT NULL COMMENT '座位数',
    `fuel_type`             VARCHAR(20)          DEFAULT NULL COMMENT '燃料类型',
    `mileage`               DECIMAL(10, 2)       DEFAULT '0.00' COMMENT '行驶里程(km)',
    `last_maintenance_date` DATE                 DEFAULT NULL COMMENT '最后维护日期',
    `next_inspection_date`  DATE                 DEFAULT NULL COMMENT '下次年检日期',
    `insurance_expiry_date` DATE                 DEFAULT NULL COMMENT '保险到期日期',
    `notes`                 TEXT COMMENT '备注',
    `created_at`            TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`            TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`            TIMESTAMP            DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`vehicle_uuid`),
    UNIQUE KEY `idx_plate_number` (`plate_number`),
    KEY `idx_status` (`status`),
    KEY `idx_model` (`model`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='车辆表';
