-- 司机资质表
-- 注意：执行此SQL前，请确保xf_driver表已经创建
CREATE TABLE IF NOT EXISTS `xf_driver_qualification`
(
    `qualification_uuid`   CHAR(36)    NOT NULL COMMENT '资质UUID',
    `driver_uuid`          CHAR(36)    NOT NULL COMMENT '司机UUID',
    `qualification_type`   VARCHAR(50) NOT NULL COMMENT '资质类型',
    `qualification_number` VARCHAR(50)          DEFAULT NULL COMMENT '资质编号',
    `issue_date`           DATE        NOT NULL COMMENT '发证日期',
    `expiry_date`          DATE        NOT NULL COMMENT '到期日期',
    `issuing_authority`    VARCHAR(100)         DEFAULT NULL COMMENT '发证机构',
    `notes`                TEXT COMMENT '备注',
    `created_at`           TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`           TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`qualification_uuid`),
    KEY `idx_driver_uuid` (`driver_uuid`),
    KEY `idx_qualification_type` (`qualification_type`),
    KEY `idx_expiry_date` (`expiry_date`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='司机资质表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_driver_qualification`
    ADD CONSTRAINT `fk_driver_qualification_driver` FOREIGN KEY (`driver_uuid`) REFERENCES `xf_driver` (`driver_uuid`)
        ON DELETE CASCADE ON UPDATE CASCADE;
