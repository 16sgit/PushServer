package push

import (
	"encoding/json"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//验证数据是否正常
func TestPushService(t *testing.T) {
	var info Info
	jsons := `{"push_type":"alipay","user":"2088022688615071","template_message_id":1,"message":{"template_id":"e4fa136c172248de9dacd0f34034971f","context":{"head_color":"#173177","url":"https:///PrimarySchoolWeChat/develop/dist/index.html#/school/AttendanceManage?identity=parent&current_child_id=1795422token=al_xbull0000001","action_name":"点击查看详情","keyword1":{"color":"#173177","value":"十牛信息校园"        },"keyword2":{"color":"#173177","value":"林国良测试学生"        },"keyword3":{"color":"#173177","value":"Two"        },"keyword4":{"color":"#173177","value":"11月20日14:27"        },"first":{"color":"#173177","value":"您好，您的小孩已打卡成功"        },"remark":{"color":"#173177","value":"详细请点击查看，感谢您的使用。"        }      }    },"config":{}}`
	//jsons := `{"user":"1"}`
	json.Unmarshal([]byte(jsons), &info)

	log.Printf("%+v", info)
}

//验证格式
func TestValidateFormat(t *testing.T) {

	Convey("验证数据格式", t, func() {
		Convey(`返回error应该为nil`, func() {
			var info Info
			jsons := `{"push_type":"alipay","user":"2088022688615071","template_message_id":1,"message":{"template_id":"e4fa136c172248de9dacd0f34034971f","context":{"head_color":"#173177","url":"https:///PrimarySchoolWeChat/develop/dist/index.html#/school/AttendanceManage?identity=parent&current_child_id=1795422token=al_xbull0000001","action_name":"点击查看详情","keyword1":{"color":"#173177","value":"十牛信息校园"        },"keyword2":{"color":"#173177","value":"林国良测试学生"        },"keyword3":{"color":"#173177","value":"Two"        },"keyword4":{"color":"#173177","value":"11月20日14:27"        },"first":{"color":"#173177","value":"您好，您的小孩已打卡成功"        },"remark":{"color":"#173177","value":"详细请点击查看，感谢您的使用。"        }      }    },"config":{}}`
			json.Unmarshal([]byte(jsons), &info)
			So(ValidateFormat(&info), ShouldBeNil)
		})
		Convey(`没有推送类型的`, func() {
			var info Info
			jsons := `{"user":"123","template_message_id":1,"message":{"template_id":"e4fa136c172248de9dacd0f34034971f","context":{"head_color":"#173177","url":"https:///PrimarySchoolWeChat/develop/dist/index.html#/school/AttendanceManage?identity=parent&current_child_id=1795422token=al_xbull0000001","action_name":"点击查看详情","keyword1":{"color":"#173177","value":"十牛信息校园"        },"keyword2":{"color":"#173177","value":"林国良测试学生"        },"keyword3":{"color":"#173177","value":"Two"        },"keyword4":{"color":"#173177","value":"11月20日14:27"        },"first":{"color":"#173177","value":"您好，您的小孩已打卡成功"        },"remark":{"color":"#173177","value":"详细请点击查看，感谢您的使用。"        }      }    },"config":{}}`
			json.Unmarshal([]byte(jsons), &info)
			err := ValidateFormat(&info)
			log.Printf("%+v", err)
			So(err, ShouldNotBeNil)
		})
		Convey(`没有模板id的`, func() {
			var info Info
			jsons := `{"push_type":"alipay","user":"2088022688615071","message":{"template_id":"e4fa136c172248de9dacd0f34034971f","context":{"head_color":"#173177","url":"https:///PrimarySchoolWeChat/develop/dist/index.html#/school/AttendanceManage?identity=parent&current_child_id=1795422token=al_xbull0000001","action_name":"点击查看详情","keyword1":{"color":"#173177","value":"十牛信息校园"        },"keyword2":{"color":"#173177","value":"林国良测试学生"        },"keyword3":{"color":"#173177","value":"Two"        },"keyword4":{"color":"#173177","value":"11月20日14:27"        },"first":{"color":"#173177","value":"您好，您的小孩已打卡成功"        },"remark":{"color":"#173177","value":"详细请点击查看，感谢您的使用。"        }      }    },"config":{}}`
			json.Unmarshal([]byte(jsons), &info)
			err := ValidateFormat(&info)
			log.Printf("%+v", err)
			So(err, ShouldNotBeNil)
		})
	})
}

//验证是否能正常获取推送服务
func TestNewPusher(t *testing.T) {
	Convey("验证获取的推送服务", t, func() {
		Convey(`获取微信推送服务`, func() {
			_, err := NewPusher("wechat")
			So(err, ShouldBeNil)
		})
		Convey(`获取不到推送服务的`, func() {
			_, err := NewPusher("aabcd")
			So(err, ShouldNotBeNil)
		})
	})
}
