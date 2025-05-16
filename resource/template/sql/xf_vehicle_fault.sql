-- 车辆故障报告表
-- 注意：执行此SQL前，请确保xf_vehicle和xf_maintenance表已经创建
CREATE TABLE IF NOT EXISTS `xf_vehicle_fault` (
  `id` varchar(36) NOT NULL COMMENT '故障ID',
  `vehicle_id` varchar(36) NOT NULL COMMENT '车辆ID',
  `reporter_id` varchar(36) NOT NULL COMMENT '报告人ID',
  `fault_type` varchar(50) NOT NULL COMMENT '故障类型',
  `fault_description` text NOT NULL COMMENT '故障描述',
  `report_date` datetime NOT NULL COMMENT '报告时间',
  `fault_location` varchar(255) DEFAULT NULL COMMENT '故障发生地点',
  `severity` tinyint(1) NOT NULL DEFAULT '2' COMMENT '严重程度: 1-轻微, 2-一般, 3-严重, 4-危险',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 0-已取消, 1-待处理, 2-处理中, 3-已解决',
  `maintenance_id` varchar(36) DEFAULT NULL COMMENT '关联的维修记录ID',
  `notes` text COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_vehicle_id` (`vehicle_id`),
  KEY `idx_reporter_id` (`reporter_id`),
  KEY `idx_report_date` (`report_date`),
  KEY `idx_status` (`status`),
  KEY `idx_maintenance_id` (`maintenance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='车辆故障报告表';

-- 外键约束（需要在所有表创建完成后执行）
-- ALTER TABLE `xf_vehicle_fault` ADD CONSTRAINT `fk_vehicle_fault_vehicle` FOREIGN KEY (`vehicle_id`) REFERENCES `xf_vehicle` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
-- ALTER TABLE `xf_vehicle_fault` ADD CONSTRAINT `fk_vehicle_fault_maintenance` FOREIGN KEY (`maintenance_id`) REFERENCES `xf_maintenance` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;
-- ALTER TABLE `xf_vehicle_fault` ADD CONSTRAINT `fk_vehicle_fault_reporter` FOREIGN KEY (`reporter_id`) REFERENCES `xf_user` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
