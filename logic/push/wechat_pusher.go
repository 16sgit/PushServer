package push

import (
	"PushServer/pkg/wechat"
	"errors"
	"reflect"

	"github.com/silenceper/wechat/template"
)

type WeChatPusher struct {
	msg    *template.Message
	mpinfo *wechat.MpInfo
}

//获取公众号配置信息
func getMpInfo() (*wechat.MpInfo,error){
	
}

//验证推送数据是否符合要求
func (p *WeChatPusher) Validate(data *Info) error {
	//获取公众号信息
	var err error
	p.msg,err := getMpInfo(&data.Config)
	
	app_id, ok := data.Config["app_id"]

	if !ok || app_id == "" {
		return errors.New("wechat:缺乏参数 appid")
	}
	p.mpinfo.AppID = app_id

	app_secret, ok := data.Config["app_secret"]
	if !ok || app_secret == "" {
		return errors.New("wechat:缺乏参数 app_secret")
	}
	p.mpinfo.AppSecret = app_secret

	token, ok := data.Config["token"]
	if ok {
		p.mpinfo.Token = token
	} else {
		p.mpinfo.Token = ""
	}

	encoding_aeskey, ok := data.Config["encoding_aeskey"]
	if ok {
		p.mpinfo.EncodingAESKey = encoding_aeskey
	} else {
		p.mpinfo.EncodingAESKey = ""
	}

	//获取模板消息
	if data.User == "" {
		return errors.New("wechat:找不到推送用户")
	}
	p.msg.ToUser = data.User

	template_id, ok := data.Message["template_id"]
	if !ok || template_id == "" {
		return errors.New("wechat:缺乏 模板id")
	}
	p.msg.TemplateID = template_id.(string)

	url, ok := data.Message["url"]
	if ok {
		p.msg.URL = url.(string)
	} else {
		p.msg.URL = ""
	}

	color, ok := data.Message["color"]
	if ok {
		p.msg.Color = color.(string)
	} else {
		p.msg.Color = ""
	}

	dataItems, ok := data.Message["data"]
	if ok && reflect.TypeOf(dataItems).String() != "map[string]interface {}" {
		return errors.New("wechat:推送的数据格式错误")
	}

	p.msg.Data = make(map[string]*template.DataItem)
	for key, dataItem := range dataItems.(map[string]interface{}) {
		item := &template.DataItem{
			Value: dataItem.(map[string]interface{})["value"].(string),
			Color: dataItem.(map[string]interface{})["color"].(string),
		}
		p.msg.Data[key] = item
	}

	return nil
}

//推送
func (p WeChatPusher) Push() (PushResponse, error) {
	return PushResponse{}, nil
}
