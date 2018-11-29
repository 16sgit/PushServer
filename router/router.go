package router

import (
	"PushServer/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerConfig.RunMode)

	//api接口
	apiRouter(r)

	return r

}
