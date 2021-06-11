package main

import "flipped-aurora/gf-vue-admin/server/boot"

// @title Gf-vue-admin Swagger Docs
// @version 1.0.0
// @description This is a Gf-vue-admin Server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	boot.Initialize()
	boot.Server.Initialize() // 初始化gf服务器
}
