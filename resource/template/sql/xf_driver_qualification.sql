-- 司机资质表
-- 注意：执行此SQL前，请确保xf_driver表已经创建
CREATE TABLE IF NOT EXISTS `xf_driver_qualification` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `driver_id` int(11) NOT NULL COMMENT '司机ID',
  `qualification_type` varchar(50) NOT NULL COMMENT '资质类型',
  `qualification_number` varchar(50) DEFAULT NULL COMMENT '资质编号',
  `issue_date` date NOT NULL COMMENT '发证日期',
  `expiry_date` date NOT NULL COMMENT '到期日期',
  `issuing_authority` varchar(100) DEFAULT NULL COMMENT '发证机构',
  `notes` text COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_driver_id` (`driver_id`),
  KEY `idx_qualification_type` (`qualification_type`),
  KEY `idx_expiry_date` (`expiry_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='司机资质表';

-- 外键约束（需要在所有表创建完成后执行）
-- ALTER TABLE `xf_driver_qualification` ADD CONSTRAINT `fk_driver_qualification_driver` FOREIGN KEY (`driver_id`) REFERENCES `xf_driver` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
