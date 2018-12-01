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
func getMpInfo(config *map[string]string) (*wechat.MpInfo, error) {
	mpinfo := wechat.MpInfo{}

	app_id, ok := (*config)["app_id"]

	if !ok || app_id == "" {
		return nil, errors.New("wechat:缺乏参数 appid")
	}
	mpinfo.AppID = app_id

	app_secret, ok := (*config)["app_secret"]
	if !ok || app_secret == "" {
		return nil, errors.New("wechat:缺乏参数 app_secret")
	}
	mpinfo.AppSecret = app_secret

	token, ok := (*config)["token"]
	if ok {
		mpinfo.Token = token
	} else {
		mpinfo.Token = ""
	}

	encoding_aeskey, ok := (*config)["encoding_aeskey"]
	if ok {
		mpinfo.EncodingAESKey = encoding_aeskey
	} else {
		mpinfo.EncodingAESKey = ""
	}

	return &mpinfo, nil
}

func getTemplate(data *Info) (*template.Message, error) {
	msg := template.Message{}

	//获取模板消息
	if data.User == "" {
		return nil, errors.New("wechat:找不到推送用户")
	}
	msg.ToUser = data.User

	template_id, ok := data.Message["template_id"]
	if !ok || template_id == "" {
		return nil, errors.New("wechat:缺乏 模板id")
	}
	msg.TemplateID = template_id.(string)

	url, ok := data.Message["url"]
	if ok {
		msg.URL = url.(string)
	} else {
		msg.URL = ""
	}

	color, ok := data.Message["color"]
	if ok {
		msg.Color = color.(string)
	} else {
		msg.Color = ""
	}

	dataItems, ok := data.Message["data"]
	if ok && reflect.TypeOf(dataItems).String() != "map[string]interface {}" {
		return nil, errors.New("wechat:推送的数据格式错误")
	}

	msg.Data = make(map[string]*template.DataItem)
	for key, dataItem := range dataItems.(map[string]interface{}) {
		item := &template.DataItem{
			Value: dataItem.(map[string]interface{})["value"].(string),
			Color: dataItem.(map[string]interface{})["color"].(string),
		}
		msg.Data[key] = item
	}

	return &msg, nil
}

//验证推送数据是否符合要求
func (p *WeChatPusher) Validate(data *Info) error {
	//获取公众号信息
	var err error
	p.mpinfo, err = getMpInfo(&data.Config)
	if err != nil {
		return err
	}

	p.msg, err = getTemplate(data)
	if err != nil {
		return err
	}

	return nil
}

//推送
func (p WeChatPusher) Push() (PushResponse, error) {
	return PushResponse{}, nil
}
