-- 排班模板表
-- 注意：执行此SQL前，请确保xf_route表已经创建
CREATE TABLE IF NOT EXISTS `xf_schedule_template` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '模板ID',
  `name` varchar(100) NOT NULL COMMENT '模板名称',
  `route_id` int(11) NOT NULL COMMENT '线路ID',
  `day_of_week` tinyint(1) NOT NULL COMMENT '星期几: 1-7 代表周一至周日',
  `start_time` time NOT NULL COMMENT '开始时间',
  `end_time` time NOT NULL COMMENT '结束时间',
  `vehicle_count` int(11) NOT NULL DEFAULT '1' COMMENT '所需车辆数',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用, 1-启用',
  `notes` text COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_route_id` (`route_id`),
  KEY `idx_day_of_week` (`day_of_week`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='排班模板表';

-- 外键约束（需要在所有表创建完成后执行）
-- ALTER TABLE `xf_schedule_template` ADD CONSTRAINT `fk_schedule_template_route` FOREIGN KEY (`route_id`) REFERENCES `xf_route` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;