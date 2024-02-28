CREATE TABLE user_device
(
    id              BIGINT(20)  PRIMARY KEY AUTO_INCREMENT      COMMENT '主键id',
    user_id         BIGINT(20)              NOT NULL            COMMENT '用户id',
    device_id       VARCHAR(20)             NOT NULL            COMMENT '设置id',
    ctime           DATETIME                NULL                COMMENT '创建时间',
    mtime           DATETIME                NULL                COMMENT '修改时间',
    deleted         BOOLEAN DEFAULT FALSE   NOT NULL            COMMENT '是否删除'
)