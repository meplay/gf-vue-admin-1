package main

import (
	"server/boot"
)

func main() {
	boot.InitializeRedis()     // 初始化Redis连接, 如果use_multipoint为false,则不会初始化redis的服务
	boot.InitializeRunServer() // 初始化gf服务器
}
