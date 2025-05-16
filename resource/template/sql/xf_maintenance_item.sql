-- 维修项目表
-- 注意：执行此SQL前，请确保xf_maintenance表已经创建
CREATE TABLE IF NOT EXISTS `xf_maintenance_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '项目ID',
  `maintenance_id` int(11) NOT NULL COMMENT '维修记录ID',
  `item_name` varchar(100) NOT NULL COMMENT '项目名称',
  `item_cost` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '项目费用',
  `item_description` text COMMENT '项目描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_maintenance_id` (`maintenance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='维修项目表';

-- 外键约束（需要在所有表创建完成后执行）
-- ALTER TABLE `xf_maintenance_item` ADD CONSTRAINT `fk_maintenance_item_maintenance` FOREIGN KEY (`maintenance_id`) REFERENCES `xf_maintenance` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;