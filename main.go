package main

import (
	"PushServer/pkg/logging"
	"PushServer/pkg/setting"
)

//初始化统一控制方法
func serverInit() {
	setting.InitSetUp("conf/app.ini")
	logging.InitSetUp()
}

//销毁统一控制
func serverDestory() {

}

func main() {
	serverInit()
	defer serverDestory()
	logging.Logger.Info("test")
}
