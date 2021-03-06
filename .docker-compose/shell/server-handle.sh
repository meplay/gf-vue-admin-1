#! /bin/bash

rm -f ./config/config.yaml
# 生成config.toml文件, 用于docker-compose的使用
touch ./config/config.yaml
filename="./config/config.yaml"
cat>"${filename}"<<EOF
# HTTP Server
server:
  address: ':8888'
  log-path: './logs/server'
  dump-router-map: true
  access-log-enabled: true
  error-log-enabled: true
  graceful: true

# Logger
logger:
  path: './logs/log'
  level: 'all'
  stdout: true

# database 配置
database:
  default: # 默认启动Mysql配置
    link: "mysql:root:gdkid,,..@tcp(177.7.0.13:3306)/gf-vue-admin"
    debug: true
  logger: # Database logger.
    path: './logs/sql'
    level: 'all'
    stdout: true

# Redis 配置
redis:
  default: '177.7.0.14:6379,1,?idleTimeout=600'
  cache: '127.0.0.1:6379,2,?idleTimeout=600'
EOF