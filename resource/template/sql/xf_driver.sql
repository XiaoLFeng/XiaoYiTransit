-- 司机表
CREATE TABLE IF NOT EXISTS `xf_driver`
(
    `driver_uuid`         CHAR(36)    NOT NULL COMMENT '司机UUID',
    `employee_id`         VARCHAR(20) NOT NULL COMMENT '工号',
    `name`                VARCHAR(50) NOT NULL COMMENT '姓名',
    `gender`              TINYINT(1)           DEFAULT NULL COMMENT '性别: 1-男, 2-女',
    `id_card`             VARCHAR(18)          DEFAULT NULL COMMENT '身份证号',
    `phone`               VARCHAR(20) NOT NULL COMMENT '联系电话',
    `emergency_contact`   VARCHAR(50)          DEFAULT NULL COMMENT '紧急联系人',
    `emergency_phone`     VARCHAR(20)          DEFAULT NULL COMMENT '紧急联系电话',
    `license_number`      VARCHAR(50) NOT NULL COMMENT '驾驶证号',
    `license_type`        VARCHAR(20) NOT NULL COMMENT '驾驶证类型',
    `license_expiry_date` DATE        NOT NULL COMMENT '驾驶证到期日期',
    `entry_date`          DATE        NOT NULL COMMENT '入职日期',
    `status`              TINYINT(1)  NOT NULL DEFAULT '1' COMMENT '状态: 0-离职, 1-在职, 2-休假, 3-停职',
    `address`             VARCHAR(255)         DEFAULT NULL COMMENT '住址',
    `notes`               TEXT COMMENT '备注',
    `created_at`          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`          TIMESTAMP            DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`driver_uuid`),
    UNIQUE KEY `idx_employee_id` (`employee_id`),
    UNIQUE KEY `idx_license_number` (`license_number`),
    KEY `idx_name` (`name`),
    KEY `idx_phone` (`phone`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='司机表';
