package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	// 读取配置
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warningln("123")
	global.Log.Error("123")
	global.DB = core.InitGorm()
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	router.Run(addr)
}
