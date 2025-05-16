-- 线路站点关联表
-- 注意：执行此SQL前，请确保xf_route和xf_station表已经创建
CREATE TABLE IF NOT EXISTS `xf_route_station` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `route_id` int(11) NOT NULL COMMENT '线路ID',
  `station_id` int(11) NOT NULL COMMENT '站点ID',
  `sequence` int(11) NOT NULL COMMENT '站点顺序',
  `distance_from_start` decimal(10,2) DEFAULT '0.00' COMMENT '距起点距离(km)',
  `estimated_time` int(11) DEFAULT NULL COMMENT '预计到达时间(分钟)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_route_station_seq` (`route_id`,`station_id`,`sequence`),
  KEY `idx_station_id` (`station_id`),
  KEY `idx_route_id` (`route_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='线路站点关联表';

-- 外键约束（需要在所有表创建完成后执行）
-- ALTER TABLE `xf_route_station` ADD CONSTRAINT `fk_route_station_route` FOREIGN KEY (`route_id`) REFERENCES `xf_route` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
-- ALTER TABLE `xf_route_station` ADD CONSTRAINT `fk_route_station_station` FOREIGN KEY (`station_id`) REFERENCES `xf_station` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
