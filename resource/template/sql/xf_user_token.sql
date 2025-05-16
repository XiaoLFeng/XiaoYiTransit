-- 用户授权Token表
-- 注意：执行此SQL前，请确保xf_user表已经创建
CREATE TABLE IF NOT EXISTS `xf_user_token`
(
    `token_uuid` CHAR(36)    NOT NULL COMMENT 'Token唯一标识UUID',
    `user_uuid`  CHAR(36)    NOT NULL COMMENT '用户UUID',
    `user_agent` VARCHAR(255)         DEFAULT NULL COMMENT '用户浏览器和设备信息',
    `client_ip`  VARCHAR(50)          DEFAULT NULL COMMENT '用户客户端IP',
    `referer`    VARCHAR(255)         DEFAULT NULL COMMENT '请求来源页面',
    `created_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `expires_at` TIMESTAMP   NOT NULL COMMENT '过期时间',
    PRIMARY KEY (`token_uuid`),
    KEY `idx_user_uuid` (`user_uuid`),
    KEY `idx_expires_at` (`expires_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户授权Token表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_user_token`
    ADD CONSTRAINT `fk_token_user` FOREIGN KEY (`user_uuid`) REFERENCES `xf_user` (`user_uuid`)
        ON DELETE CASCADE ON UPDATE CASCADE;
