package push

import (
	"PushServer/pkg/setting"
	"encoding/json"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Init() {
	setting.InitSetUp("../../conf/app.ini")
}

//微信模板推送数据验证
func TestWechatValidate(t *testing.T) {
	Convey("微信模板推送数据验证", t, func() {
		Convey(`正常，没有任何报错的`, func() {
			var info Info
			jsons := `{"push_type":"wechat_template","user":"ofCONv1WVGN2zjF-lOv-rF_nuTc8","template_message_id":1,"message":{"template_id":"D4thuhB44AWUzoP-GcV9IegoCeCycbm41BafOHGY3F8","url":"http://www/baidu.com","data":{"keyword1":{"color":"#173177","value":"十牛信息校园"}}},"config":{"app_id":"wx4b5e49b637c4ccf1","app_secret":"f92fe88d88bcf7de81f1b069fcb166cf"}}`
			json.Unmarshal([]byte(jsons), &info)
			wx := &WeChatPusher{
				msg:    &template.Message{},
				mpinfo: &wechat.MpInfo{},
			}
			So(wx.Validate(&info), ShouldBeNil)
		})
		Convey(`缺乏app_id的`, func() {
			var info Info
			jsons := `{"push_type":"wechat_template","user":"ofCONv1WVGN2zjF-lOv-rF_nuTc8","template_message_id":1,"message":{"template_id":"D4thuhB44AWUzoP-GcV9IegoCeCycbm41BafOHGY3F8","url":"http://www/baidu.com","data":{"keyword1":{"color":"#173177","value":"十牛信息校园"}}},"config":{"app_id":"","app_secret":"f92fe88d88bcf7de81f1b069fcb166cf"}}`
			json.Unmarshal([]byte(jsons), &info)
			wx := &WeChatPusher{
				msg:    &template.Message{},
				mpinfo: &wechat.MpInfo{},
			}
			err := wx.Validate(&info)
			log.Println(err)
			So(err, ShouldNotBeNil)
		})
		Convey(`缺乏app_secret的`, func() {
			var info Info
			jsons := `{"push_type":"wechat_template","user":"ofCONv1WVGN2zjF-lOv-rF_nuTc8","template_message_id":1,"message":{"template_id":"D4thuhB44AWUzoP-GcV9IegoCeCycbm41BafOHGY3F8","url":"http://www/baidu.com","data":{"keyword1":{"color":"#173177","value":"十牛信息校园"}}},"config":{"app_id":"wx4b5e49b637c4ccf1","app_secret":""}}`
			json.Unmarshal([]byte(jsons), &info)
			wx := &WeChatPusher{
				msg:    &template.Message{},
				mpinfo: &wechat.MpInfo{},
			}
			err := wx.Validate(&info)
			log.Println(err)
			So(err, ShouldNotBeNil)
		})
		Convey(`推送数据格式错误的`, func() {
			var info Info
			jsons := `{"push_type":"wechat_template","user":"ofCONv1WVGN2zjF-lOv-rF_nuTc8","template_message_id":1,"message":{"template_id":"D4thuhB44AWUzoP-GcV9IegoCeCycbm41BafOHGY3F8","url":"http://www/baidu.com","data":"abc"},"config":{"app_id":"wx4b5e49b637c4ccf1","app_secret":"f92fe88d88bcf7de81f1b069fcb166cf"}}`
			json.Unmarshal([]byte(jsons), &info)
			wx := &WeChatPusher{
				msg:    &template.Message{},
				mpinfo: &wechat.MpInfo{},
			}
			err := wx.Validate(&info)
			log.Println(err)
			So(err, ShouldNotBeNil)
		})

	})
}

//模板推送测试
func TestWechatPush(t *testing.T) {
	Init()

	Convey("微信模板推送", t, func() {
		Convey(`正常，没有任何报错的`, func() {
			var info Info
			jsons := `{"push_type":"wechat","user":"ofCONv1WVGN2zjF-lOv-rF_nuTc8","template_message_id":1,"message":{"template_id":"D4thuhB44AWUzoP-GcV9IegoCeCycbm41BafOHGY3F8","url":"http://www/baidu.com","data":{"keyword1":{"color":"#173177","value":"十牛信息校园"}}},"config":{"app_id":"wx4b5e49b637c4ccf1","app_secret":"f92fe88d88bcf7de81f1b069fcb166cf"}}`
			json.Unmarshal([]byte(jsons), &info)
			wx := &WeChatTemplatePusher{}
			err := wx.Validate(&info)
			if err != nil {
				log.Println(err)
			}
			response, _ := wx.Push()
			log.Printf("%+v", response)
		})
	})
}
