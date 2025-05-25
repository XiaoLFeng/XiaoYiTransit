-- 车辆保险记录表
CREATE TABLE IF NOT EXISTS `xf_vehicle_insurance`
(
    `insurance_uuid`    CHAR(36)    NOT NULL COMMENT '保险UUID',
    `vehicle_uuid`      CHAR(36)    NOT NULL COMMENT '车辆UUID',
    `insurance_type`    VARCHAR(50) NOT NULL COMMENT '保险类型',
    `policy_number`     VARCHAR(50) NOT NULL COMMENT '保单号',
    `insurer`           VARCHAR(100)         DEFAULT NULL COMMENT '保险公司',
    `start_date`        DATE        NOT NULL COMMENT '保险开始日期',
    `expiry_date`       DATE        NOT NULL COMMENT '保险到期日期',
    `coverage_amount`   DECIMAL(12, 2)       DEFAULT '0.00' COMMENT '保险金额',
    `premium`           DECIMAL(10, 2)       DEFAULT '0.00' COMMENT '保费',
    `payment_date`      DATE                 DEFAULT NULL COMMENT '缴费日期',
    `payment_method`    VARCHAR(50)          DEFAULT NULL COMMENT '缴费方式',
    `contact_person`    VARCHAR(50)          DEFAULT NULL COMMENT '联系人',
    `contact_phone`     VARCHAR(20)          DEFAULT NULL COMMENT '联系电话',
    `notes`             TEXT COMMENT '备注',
    `created_at`        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        TIMESTAMP            DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`insurance_uuid`),
    KEY `idx_vehicle_uuid` (`vehicle_uuid`),
    KEY `idx_policy_number` (`policy_number`),
    KEY `idx_start_date` (`start_date`),
    KEY `idx_expiry_date` (`expiry_date`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='车辆保险记录表';