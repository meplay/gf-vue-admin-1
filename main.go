package main

import "github.com/flipped-aurora/gf-vue-admin/boot"

func main() {
	boot.Initialize()
	boot.Server.Initialize() // 初始化gf服务器
}
