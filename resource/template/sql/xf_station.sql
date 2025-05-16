-- 站点表
CREATE TABLE IF NOT EXISTS `xf_station` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '站点ID',
  `name` varchar(100) NOT NULL COMMENT '站点名称',
  `code` varchar(20) DEFAULT NULL COMMENT '站点编码',
  `address` varchar(255) DEFAULT NULL COMMENT '站点地址',
  `longitude` decimal(10,6) DEFAULT NULL COMMENT '经度',
  `latitude` decimal(10,6) DEFAULT NULL COMMENT '纬度',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 0-停用, 1-启用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='站点表';