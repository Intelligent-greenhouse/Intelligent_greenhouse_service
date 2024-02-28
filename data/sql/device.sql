CREATE TABLE device
(
    id              BIGINT(20)  PRIMARY KEY AUTO_INCREMENT      COMMENT '主键id',
    device_id       TEXT                    NOT NULL            COMMENT '设备码',
    is_activation   BOOLEAN                 NOT NULL            COMMENT '激活状态',
    ctime           DATETIME                NULL                COMMENT '创建时间',
    mtime           DATETIME                NULL                COMMENT '修改时间',
    deleted         BOOLEAN DEFAULT FALSE   NOT NULL            COMMENT '是否删除'
)