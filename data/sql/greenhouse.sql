CREATE TABLE greenhouse
(
    id              BIGINT(20)  PRIMARY KEY AUTO_INCREMENT      COMMENT '主键id',
    name            TEXT                    NOT NULL            COMMENT '大棚名称',
    pos             TEXT                    NULL                COMMENT '位置',
    size            BIGINT(20)              NULL                COMMENT '大小',
    des             TEXT                    NULL                COMMENT '备注',
    ctime           DATETIME                NULL                COMMENT '创建时间',
    mtime           DATETIME                NULL                COMMENT '修改时间',
    deleted         BOOLEAN DEFAULT FALSE   NOT NULL            COMMENT '是否删除'
)