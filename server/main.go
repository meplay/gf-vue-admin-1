package main

import "gf-vue-admin/boot"

func main() {
	boot.Config.Initialize() // 初始化配置
	boot.Server.Initialize() // 初始化gf服务器
}
