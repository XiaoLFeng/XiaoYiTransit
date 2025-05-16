-- 维修项目表
CREATE TABLE IF NOT EXISTS `xf_maintenance_item`
(
    `maintenance_item_uuid` CHAR(36)       NOT NULL COMMENT '项目UUID',
    `maintenance_uuid`      CHAR(36)       NOT NULL COMMENT '维修记录UUID',
    `item_name`             VARCHAR(100)   NOT NULL COMMENT '项目名称',
    `item_cost`             DECIMAL(10, 2) NOT NULL DEFAULT '0.00' COMMENT '项目费用',
    `item_description`      TEXT COMMENT '项目描述',
    `created_at`            TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`            TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`maintenance_item_uuid`),
    KEY `idx_maintenance_uuid` (`maintenance_uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='维修项目表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_maintenance_item`
    ADD CONSTRAINT `fk_maintenance_item_maintenance` FOREIGN KEY (`maintenance_uuid`) REFERENCES `xf_maintenance` (`maintenance_uuid`)
        ON DELETE CASCADE ON UPDATE CASCADE;