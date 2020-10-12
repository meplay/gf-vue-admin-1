package main

import (
	"server/boot"
)

func main() {
	// 如需自动导入初始数据表及数据,打开以下注释内容,初始化完成后请删掉或注释以下内容
	// boot.InitializeDataTableAndData() // 初始化数据表与数据
	boot.InitializeI18n()      // 初始化I18N国际化
	boot.InitializeRedis()     // 初始化Redis连接, 如果use_multipoint为false,则不会初始化redis的服务
	boot.InitializeRunServer() // 初始化gf服务器
}
