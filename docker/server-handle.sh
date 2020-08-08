#! /bin/bash

rm -f ./boot/server.go
# 生成server.go文件, 添加Router.Static("/admin", "./resource/dist")这个代码
touch ./boot/server.go
filename="./boot/server.go"
cat>"${filename}"<<EOF
package boot

import (
	"server/app/middleware"
	"server/router"
	"time"

	"github.com/gogf/gf/frame/g"
)

func InitializeRunServer() {
	s := g.Server()
	s.Use(middleware.Error)
	s.SetReadTimeout(10 * time.Second)
	s.SetWriteTimeout(10 * time.Second)
	s.SetMaxHeaderBytes(1 << 20)
	s.SetIndexFolder(true)
	s.AddStaticPath("/form-generator", "public/page")
	s.AddStaticPath("/admin", "public/dist")
	router.InitializeRouters()
	s.Run()
}
EOF

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

# system configuration
[system]
    Db = "default"
    UseMultipoint = true
    Env = "public" # Change to "develop" to skip authentication for development mode

# 请自行七牛申请对应的 公钥 私钥 bucket 和 域名地址
[qiniu]
    AccessKey   = ""
    SecretKey   = ""
    Bucket      = ""
    ImgPath     = ""

# captcha configuration
[captcha]
    KeyLong     = 6
    ImgWidth    = 240
    ImgHeight   = 80

# casbin configuration
[casbin]
    ModelPath =  "./public/rbac_model.conf"

# jwt configuration
[jwt]
    SigningKey  = "SliverHorn"
    ExpiresAt = 1 # 3600000000000秒 == 1天 默认设置为1天
    RefreshAt = 168 # 3600000000000秒 == 1天 , 24 * 7 = 168 刷新的token设置为一星期

# Logger
[logger]
    Path        = "./logs/log"
    Level       = "all"
    Stdout      = true

# MySQL 配置
[database]
    # 默认启动Mysql配置
    [database.default]
    Link      = "mysql:root:gdkid,,..@tcp(gfmysql)/gf-vue-admin"
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