-- 维修记录表
-- 注意：执行此SQL前，请确保xf_vehicle表已经创建
CREATE TABLE IF NOT EXISTS `xf_maintenance` (
  `id` varchar(36) NOT NULL COMMENT '维修ID',
  `vehicle_id` varchar(36) NOT NULL COMMENT '车辆ID',
  `maintenance_type` tinyint(1) NOT NULL COMMENT '维修类型: 1-常规保养, 2-故障维修, 3-事故维修, 4-年检维修',
  `description` text NOT NULL COMMENT '维修描述',
  `maintenance_date` date NOT NULL COMMENT '维修日期',
  `completion_date` date DEFAULT NULL COMMENT '完成日期',
  `cost` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '维修费用',
  `mileage` decimal(10,2) DEFAULT NULL COMMENT '维修时里程数',
  `maintenance_location` varchar(255) DEFAULT NULL COMMENT '维修地点',
  `maintenance_staff` varchar(50) DEFAULT NULL COMMENT '维修人员',
  `parts_replaced` text COMMENT '更换的零部件',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 0-已取消, 1-待维修, 2-维修中, 3-已完成',
  `notes` text COMMENT '备注',
  `created_by` varchar(36) DEFAULT NULL COMMENT '创建人ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_vehicle_id` (`vehicle_id`),
  KEY `idx_maintenance_date` (`maintenance_date`),
  KEY `idx_maintenance_type` (`maintenance_type`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='维修记录表';

-- 外键约束（需要在所有表创建完成后执行）
-- ALTER TABLE `xf_maintenance` ADD CONSTRAINT `fk_maintenance_vehicle` FOREIGN KEY (`vehicle_id`) REFERENCES `xf_vehicle` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
