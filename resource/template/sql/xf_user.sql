-- 用户表
CREATE TABLE IF NOT EXISTS `xf_user`
(
    `user_uuid`       CHAR(36)     NOT NULL COMMENT '用户ID',
    `username`        VARCHAR(50)  NOT NULL COMMENT '用户名',
    `password`        VARCHAR(100) NOT NULL COMMENT '密码',
    `real_name`       VARCHAR(50)           DEFAULT NULL COMMENT '真实姓名',
    `email`           VARCHAR(100)          DEFAULT NULL COMMENT '邮箱',
    `phone`           VARCHAR(20)           DEFAULT NULL COMMENT '手机号',
    `avatar`          VARCHAR(255)          DEFAULT NULL COMMENT '头像',
    `role_uuid`       CHAR(36)     NOT NULL COMMENT '角色ID',
    `status`          TINYINT(1)   NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用, 1-启用',
    `last_login_time` TIMESTAMP             DEFAULT NULL COMMENT '最后登录时间',
    `created_at`      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`      TIMESTAMP             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`user_uuid`),
    UNIQUE KEY `idx_username` (`username`),
    KEY `idx_email` (`email`),
    KEY `idx_phone` (`phone`),
    KEY `idx_role_id` (`role_uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

-- 外键约束（需要在所有表创建完成后执行）
ALTER TABLE `xf_user`
    ADD CONSTRAINT `fk_user_role` FOREIGN KEY (`role_uuid`) REFERENCES `xf_role` (`role_uuid`)
        ON DELETE RESTRICT ON UPDATE CASCADE;
