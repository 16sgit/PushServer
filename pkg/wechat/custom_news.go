package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/util"
)

const (
	custom_message_url = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

//文字内容
type Article struct {
	Title       string `xml:"Title,omitempty"json:"title"`
	Description string `xml:"Description,omitempty"json:"description"`
	PicURL      string `xml:"PicUrl,omitempty"json:"picUrl"`
	URL         string `xml:"Url,omitempty"json:"url"`
}

//微信客服消息 图文消息
type NewsMessage struct {
	ToUser  string                `json:"touser"`
	MsgType string                `json:"msgtype"`
	News    map[string][]*Article `json:"news"`
}

//客服
type CustomService struct {
	wx *wechat.Wechat
}

//发送图文消息
func (s CustomService) SendNews(news *NewsMessage) (rest util.CommonError, err error) {
	access_token, err := s.wx.GetAccessToken()
	if err != nil {
		return util.CommonError{}, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", custom_message_url, access_token)
	response, err := util.PostJSON(uri, news)

	err = json.Unmarshal(response, &rest)

	return rest, nil
}
