-- 角色表
CREATE TABLE IF NOT EXISTS `xf_role`
(
    `role_uuid`   CHAR(36)    NOT NULL COMMENT '角色UUID',
    `name`        VARCHAR(50) NOT NULL COMMENT '角色名称',
    `code`        VARCHAR(50) NOT NULL COMMENT '角色编码',
    `description` VARCHAR(255)         DEFAULT NULL COMMENT '角色描述',
    `status`      TINYINT(1)  NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用, 1-启用',
    `created_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  TIMESTAMP            DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`role_uuid`),
    UNIQUE KEY `idx_code` (`code`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='角色表';
