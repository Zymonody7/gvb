package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warningln("123")
	global.Log.Error("123")
	global.DB = core.InitGorm()
}
