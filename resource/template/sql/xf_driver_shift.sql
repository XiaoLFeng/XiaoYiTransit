-- 司机排班表
-- 注意：执行此SQL前，请确保xf_driver表已经创建
CREATE TABLE IF NOT EXISTS `xf_driver_shift` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `driver_id` int(11) NOT NULL COMMENT '司机ID',
  `shift_date` date NOT NULL COMMENT '排班日期',
  `shift_type` tinyint(1) NOT NULL COMMENT '班次类型: 1-早班, 2-中班, 3-晚班, 4-全天班',
  `start_time` time NOT NULL COMMENT '开始时间',
  `end_time` time NOT NULL COMMENT '结束时间',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 0-取消, 1-待执行, 2-执行中, 3-已完成',
  `notes` text COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_driver_id` (`driver_id`),
  KEY `idx_shift_date` (`shift_date`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='司机排班表';

-- 外键约束（需要在所有表创建完成后执行）
-- ALTER TABLE `xf_driver_shift` ADD CONSTRAINT `fk_driver_shift_driver` FOREIGN KEY (`driver_id`) REFERENCES `xf_driver` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;