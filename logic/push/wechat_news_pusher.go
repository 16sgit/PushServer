//微信客服消息推送功能
package push

import (
	"PushServer/pkg/wechat"

	"github.com/silenceper/wechat/util"

	"errors"
)

type WeChatNewsPusher struct {
	article       *wechat.Article
	customService *wechat.CustomService
	toUser        string
}

//获取推送的模板信息
func getArticle(data map[string]interface{}) (*wechat.Article, error) {
	article := wechat.Article{}

	title, ok := data["title"]
	if !ok || title == "" {
		return nil, errors.New("wechat:缺乏 内容标题")
	}
	article.Title = title.(string)

	description, ok := data["description"]
	if !ok || description == "" {
		return nil, errors.New("wechat:缺乏 内容描述")
	}
	article.Description = description.(string)

	picUrl, ok := data["picUrl"]
	if !ok || picUrl == "" {
		return nil, errors.New("wechat:缺乏 封面链接")
	}
	article.PicURL = picUrl.(string)

	url, ok := data["url"]
	if !ok || url == "" {
		return nil, errors.New("wechat:缺乏 跳转链接")
	}
	article.URL = url.(string)

	return &article, nil
}

//验证推送数据是否符合要求
func (p *WeChatNewsPusher) Validate(data *Info) error {
	//获取公众号信息
	var err error
	mpinfo, err := getMpInfo(&data.Config)
	if err != nil {
		return err
	}
	p.customService = mpinfo.GetCustomService()

	p.article, err = getArticle(data.Message)
	if err != nil {
		return err
	}

	if data.User == "" {
		return errors.New("wechat:找不到推送用户")
	}
	p.toUser = data.User

	return nil
}

//推送
func (p WeChatNewsPusher) Push() (response util.CommonError, err error) {
	articles := make([]*wechat.Article, 1)
	articles[0] = p.article

	news_message := wechat.NewsMessage{
		ToUser:  p.toUser,
		MsgType: "news",
		News: map[string][]*wechat.Article{
			"articles": articles,
		},
	}

	return p.customService.SendNews(&news_message)
}
