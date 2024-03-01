CREATE TABLE device
(
    id                      BIGINT(20)  PRIMARY KEY AUTO_INCREMENT      COMMENT '主键id',
    device_id               TEXT                    NOT NULL            COMMENT '设备码',

    co2                     FLOAT(20)               NULL                COMMENT '二氧化碳浓度',
    light_intensity         FLOAT(20)               NULL                COMMENT '光照强度',
    air_temperature         FLOAT(20)               NULL                COMMENT '空气温度',
    air_humidity            FLOAT(20)               NULL                COMMENT '空气湿度',
    soil_temperature        FLOAT(20)               NULL                COMMENT '土壤温度',
    soil_moisture           FLOAT(20)               NULL                COMMENT '土壤水分',
    soil_conductivity       FLOAT(20)               NULL                COMMENT '土壤导电率',
    soil_ph                 FLOAT(20)               NULL                COMMENT '土壤ph值',

    led                     BOOLEAN                 NULL                COMMENT'led开关',
    fan                     BOOLEAN                 NULL                COMMENT'风扇开关',
    water                   BOOLEAN                 NULL                COMMENT'水阀开关',
    chemical_fertilizer     BOOLEAN                 NULL                COMMENT'施肥开关',
    increase_temperature    BOOLEAN                 NULL                COMMENT'增加温度开关',
    reduce_temperature      BOOLEAN                 NULL                COMMENT'减少温度开关',
    buzzer                  BOOLEAN                 NULL                COMMENT'蜂鸣器开关',

    is_activation           BOOLEAN                 NOT NULL            COMMENT '激活状态',
    run_time                DATETIME                NULL                COMMENT '正常运行时间戳',
    ctime                   DATETIME                NULL                COMMENT '创建时间',
    mtime                   DATETIME                NULL                COMMENT '修改时间',
    del                     BOOLEAN DEFAULT FALSE   NOT NULL            COMMENT '是否删除'
)