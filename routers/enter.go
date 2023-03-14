package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouter := router.Group("/api")
	routerGroupApp := RouterGroup{apiRouter}
	routerGroupApp.SettingsRouter()
	return router
}
