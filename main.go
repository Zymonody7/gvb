package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	core.InitConf()
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
