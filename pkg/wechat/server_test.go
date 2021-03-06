package wechat

import (
	"PushServer/pkg/setting"
	"log"
	"testing"
	//	"github.com/silenceper/wechat"
	//	"github.com/silenceper/wechat/template"
	//	. "github.com/smartystreets/goconvey/convey"
)

func Init() {
	setting.InitSetUp("../../conf/app.ini")
}

//验证是否能正常获取微信操作对象
//func TestGetWechat(t *testing.T) {
//	Convey("验证是否能正常获取微信操作对象", t, func() {
//		Convey(`正常获取`, func() {
//			info := MpInfo{
//				AppID:          "wx4b5e49b637c4ccf1",
//				AppSecret:      "f92fe88d88bcf7de81f1b069fcb166cf",
//				Token:          "",
//				EncodingAESKey: "",
//			}

//			So(GetWechat(&info), ShouldHaveSameTypeAs, &wechat.Wechat{})
//		})
//		Convey(`正常获取之后，可以直接从map获取值`, func() {
//			info := MpInfo{
//				AppID:          "wx4b5e49b637c4ccf1",
//				AppSecret:      "f92fe88d88bcf7de81f1b069fcb166cf",
//				Token:          "",
//				EncodingAESKey: "",
//			}

//			So(GetWechat(&info), ShouldHaveSameTypeAs, wechats[info.AppID])
//		})
//		Convey(`刚开始应该是不存在这个值得`, func() {
//			info := MpInfo{
//				AppID:          "aaa",
//				AppSecret:      "f92fe88d88bcf7de81f1b069fcb166cf",
//				Token:          "",
//				EncodingAESKey: "",
//			}

//			_, ok := wechats[info.AppID]
//			So(ok, ShouldBeFalse)
//		})
//	})
//}

//验证是否能正常发送模板消息
//func TestPushTemplateMessage(t *testing.T) {
//	Init()

//	Convey("验证是否能正常发送模板消息", t, func() {
//		info := MpInfo{
//			AppID:          "wx4b5e49b637c4ccf1",
//			AppSecret:      "f92fe88d88bcf7de81f1b069fcb166cf",
//			Token:          "",
//			EncodingAESKey: "",
//		}

//		msg := &template.Message{
//			ToUser:     "ofCONv1WVGN2zjF-lOv-rF_nuTc8",
//			TemplateID: "D4thuhB44AWUzoP-GcV9IegoCeCycbm41BafOHGY3F8",
//			URL:        "http://www/baidu.com",
//			Color:      "",
//			Data:       map[string]*template.DataItem{},
//		}

//		wc := GetWechat(&info)
//		send_handle := wc.GetTemplate()
//		_, err := send_handle.Send(msg)

//		So(err, ShouldBeNil)
//	})
//}

//验证能否正常发送客服消息
//func TestSendCustomNewsMessage(t *testing.T) {
//	Init()

//	articles := make([]*Article, 1)
//	articles[0] = &Article{
//		Title:       "你好",
//		Description: "这个是描述",
//		PicURL:      "https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=893776491,1251327685&fm=58",
//		URL:         "http://www.baidu.com",
//	}

//	news_message := NewsMessage{
//		ToUser:  "ofCONv1WVGN2zjF-lOv-rF_nuTc8",
//		MsgType: "news",
//		News: map[string][]*Article{
//			"articles": articles,
//		},
//	}

//	info := MpInfo{
//		AppID:          "wx4b5e49b637c4ccf1",
//		AppSecret:      "f92fe88d88bcf7de81f1b069fcb166cf",
//		Token:          "",
//		EncodingAESKey: "",
//	}

//	custom_service := info.GetCustomService()

//	log.Println(custom_service.SendNews(&news_message))
//}
