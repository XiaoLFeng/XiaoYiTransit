-- 调度表
-- 注意：执行此SQL前，请确保xf_route、xf_vehicle和xf_driver表已经创建
CREATE TABLE IF NOT EXISTS `xf_schedule` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '调度ID',
  `route_id` int(11) NOT NULL COMMENT '线路ID',
  `vehicle_id` int(11) NOT NULL COMMENT '车辆ID',
  `driver_id` int(11) NOT NULL COMMENT '司机ID',
  `schedule_date` date NOT NULL COMMENT '调度日期',
  `start_time` time NOT NULL COMMENT '开始时间',
  `end_time` time NOT NULL COMMENT '结束时间',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 0-取消, 1-待执行, 2-执行中, 3-已完成',
  `actual_start_time` time DEFAULT NULL COMMENT '实际开始时间',
  `actual_end_time` time DEFAULT NULL COMMENT '实际结束时间',
  `mileage` decimal(10,2) DEFAULT '0.00' COMMENT '行驶里程(km)',
  `fuel_consumption` decimal(10,2) DEFAULT '0.00' COMMENT '油耗(L)',
  `passenger_count` int(11) DEFAULT '0' COMMENT '载客人数',
  `notes` text COMMENT '备注',
  `created_by` int(11) DEFAULT NULL COMMENT '创建人ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_route_id` (`route_id`),
  KEY `idx_vehicle_id` (`vehicle_id`),
  KEY `idx_driver_id` (`driver_id`),
  KEY `idx_schedule_date` (`schedule_date`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='调度表';

-- 外键约束（需要在所有表创建完成后执行）
-- ALTER TABLE `xf_schedule` ADD CONSTRAINT `fk_schedule_route` FOREIGN KEY (`route_id`) REFERENCES `xf_route` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
-- ALTER TABLE `xf_schedule` ADD CONSTRAINT `fk_schedule_vehicle` FOREIGN KEY (`vehicle_id`) REFERENCES `xf_vehicle` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
-- ALTER TABLE `xf_schedule` ADD CONSTRAINT `fk_schedule_driver` FOREIGN KEY (`driver_id`) REFERENCES `xf_driver` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
