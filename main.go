package main

import (
	"PushServer/logic"
	"PushServer/models"
	"PushServer/pkg/cache"
	"PushServer/pkg/logging"
	"PushServer/pkg/setting"
	//"PushServer/router"
	//"fmt"
	//"net/http"
)

//初始化统一控制方法
func serverInit() {
	setting.InitSetUp("conf/app.ini")
	logging.InitSetUp()
	models.InitSetUp()
	cache.InitSetUp()
}

//销毁统一控制
func serverDestory() {
	models.CloseDb()
}

func main() {
	serverInit()
	defer serverDestory()
	logic.Push_service()
	//	s := &http.Server{
	//		Addr:           fmt.Sprintf(":%d", setting.ServerConfig.HttpPort),
	//		Handler:        router.InitRouter(),
	//		ReadTimeout:    setting.ServerConfig.ReadTimeout,
	//		WriteTimeout:   setting.ServerConfig.WriteTimeout,
	//		MaxHeaderBytes: 1 << 20,
	//	}

	//	s.ListenAndServe()
}
