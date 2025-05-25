-- 系统表
CREATE TABLE `xf_system`
(
    `system_uuid` CHAR(36)     NOT NULL COMMENT '系统UUID',
    `key`         VARCHAR(255) NOT NULL COMMENT '系统键',
    `val`         TEXT COMMENT '系统值',
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`system_uuid`),
    UNIQUE KEY `idx_key` (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='系统表';