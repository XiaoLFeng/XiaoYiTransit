-- 司机排班表
-- 注意：执行此SQL前，请确保xf_driver表已经创建
CREATE TABLE IF NOT EXISTS `xf_driver_shift`
(
    `shift_uuid`  CHAR(36)   NOT NULL COMMENT '排班UUID',
    `driver_uuid` CHAR(36)   NOT NULL COMMENT '司机UUID',
    `shift_date`  DATE       NOT NULL COMMENT '排班日期',
    `shift_type`  TINYINT(1) NOT NULL COMMENT '班次类型: 1-早班, 2-中班, 3-晚班, 4-全天班',
    `start_time`  TIME       NOT NULL COMMENT '开始时间',
    `end_time`    TIME       NOT NULL COMMENT '结束时间',
    `status`      TINYINT(1) NOT NULL DEFAULT '1' COMMENT '状态: 0-取消, 1-待执行, 2-执行中, 3-已完成',
    `notes`       TEXT COMMENT '备注',
    `created_at`  TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`shift_uuid`),
    KEY `idx_driver_uuid` (`driver_uuid`),
    KEY `idx_shift_date` (`shift_date`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='司机排班表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_driver_shift`
    ADD CONSTRAINT `fk_driver_shift_driver` FOREIGN KEY (`driver_uuid`) REFERENCES `xf_driver` (`driver_uuid`)
        ON DELETE CASCADE ON UPDATE CASCADE;