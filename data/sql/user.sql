CREATE TABLE user
(
    id              BIGINT(20)  PRIMARY KEY AUTO_INCREMENT      COMMENT '主键id',
    username        VARCHAR(50)             NOT NULL            COMMENT '用户名',
    password        VARCHAR(20)             NOT NULL            COMMENT '密码',
    is_admin        BOOLEAN                 NOT NULL            COMMENT '是否为管理员',
    ctime           DATETIME                NULL                COMMENT '创建时间',
    mtime           DATETIME                NULL                COMMENT '修改时间',
    deleted         BOOLEAN DEFAULT FALSE   NOT NULL            COMMENT '是否删除'
)