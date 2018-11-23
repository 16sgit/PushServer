package logging

import (
	"PushServer/pkg/setting"
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
)

func InitTest() {
	setting.InitSetUp("../../conf/app.ini")

}

func TestLog(t *testing.T) {
	InitTest()
	Test()
}
