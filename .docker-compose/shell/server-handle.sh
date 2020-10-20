#! /bin/bash

rm -f ./config/config.toml
# 生成config.toml文件, 用于docker-compose的使用
touch ./config/config.toml
filename="./config/config.toml"
cat>"${filename}"<<EOF
# HTTP Server
[server]
    Address          = ":8888"
    ServerRoot       = ""
    ServerAgent      = ""
    IndexFiles       = []
    AccessLogEnabled = true
    ErrorLogEnabled  = true
    PProfEnabled     = false
    LogPath          = "./logs/server"
    SessionIdName    = ""
    SessionPath      = ""
    SessionMaxAge    = ""
    DumpRouterMap    = true

# Logger
[logger]
    Path        = "./logs/log"
    Level       = "all"
    Stdout      = true

# MySQL 配置
[database]
    # 默认启动Mysql配置
    [database.default]
    Link      = "mysql:root:gdkid,,..@tcp(mysql)/gf-vue-admin"
    Debug     = true

    # Database logger.
    [database.logger]
        Path   = "./logs/sql"
        Level  = "all"
        Stdout = true

# Redis 配置
[redis]
    # host:port[,db,pass?maxIdle=x&maxActive=x&idleTimeout=x&maxConnLifetime=x]
    default = "redis:6379,1,?idleTimeout=600"
    cache   = "127.0.0.1:6379,2,?idleTimeout=600"
EOF