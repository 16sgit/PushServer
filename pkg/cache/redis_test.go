//推送服务测试
package cache

import (
	"PushServer/models"
	"PushServer/pkg/setting"
	"log"
	"testing"
	//	"github.com/Unknwon/com"
	//. "github.com/smartystreets/goconvey/convey"
)

func Init() {
	setting.InitSetUp("../../conf/app.ini")
	models.InitSetUp()
	InitSetUp()
}

//销毁统一控制
func serverDestory() {
	models.CloseDb()
}

//模板推送功能测试
func TestWechatValidate(t *testing.T) {
	Init()
	defer serverDestory()

	log.Println(Pop())

}
