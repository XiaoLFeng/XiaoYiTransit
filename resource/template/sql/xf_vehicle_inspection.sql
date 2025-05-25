-- 车辆年检记录表
CREATE TABLE IF NOT EXISTS `xf_vehicle_inspection`
(
    `inspection_uuid`   CHAR(36)    NOT NULL COMMENT '年检UUID',
    `vehicle_uuid`      CHAR(36)    NOT NULL COMMENT '车辆UUID',
    `inspection_date`   DATE        NOT NULL COMMENT '年检日期',
    `expiry_date`       DATE        NOT NULL COMMENT '有效期截止日期',
    `inspection_result` TINYINT(1)  NOT NULL DEFAULT '1' COMMENT '年检结果: 1-通过, 2-不通过',
    `inspection_agency` VARCHAR(100)         DEFAULT NULL COMMENT '检测机构',
    `inspector`         VARCHAR(50)          DEFAULT NULL COMMENT '检测人员',
    `certificate_no`    VARCHAR(50)          DEFAULT NULL COMMENT '检测证书编号',
    `cost`              DECIMAL(10, 2)       DEFAULT '0.00' COMMENT '年检费用',
    `notes`             TEXT COMMENT '备注',
    `created_at`        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        TIMESTAMP            DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`inspection_uuid`),
    KEY `idx_vehicle_uuid` (`vehicle_uuid`),
    KEY `idx_inspection_date` (`inspection_date`),
    KEY `idx_expiry_date` (`expiry_date`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='车辆年检记录表';