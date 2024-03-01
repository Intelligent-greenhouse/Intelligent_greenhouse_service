CREATE TABLE greenhouse
(
    id              BIGINT(20)  PRIMARY KEY AUTO_INCREMENT      COMMENT '主键id',
    pos             TEXT                    NULL                COMMENT '位置',
    size            BIGINT(20)              NULL                COMMENT '大小',
    des             TEXT                    NULL                COMMENT '备注',
    ctime           DATETIME                NULL                COMMENT '创建时间',
    mtime           DATETIME                NULL                COMMENT '修改时间',
    del             BOOLEAN DEFAULT FALSE   NOT NULL            COMMENT '是否删除'
)