package push

import (
	"errors"
)

//错误类型
var PUSHER_NOT_EXIST = errors.New("未找到推送服务")

//消息体
type Info struct {
	User              string                 `json:"user"`                //推送用户
	TemplateMessageId int64                  `json:"template_message_id"` //日志id
	Config            map[string]string      `json:"config"`              //存储相关的配置信息
	Message           map[string]interface{} `json:"message"`             //存储要推送的信息
	PushType          string                 `json:"push_type"`           //推送类型，例如：wechat
}

//推送服务响应
type PushResponse struct {
	Status  int
	Message string
	Msgid   string
}

//验证消息体的格式是否正确
func ValidateFormat(info *Info) error {
	switch true {
	case info.User == "":
		return errors.New("推送用户不能为空")
		break
	case info.TemplateMessageId <= 0:
		return errors.New("模板日志id不能为空")
		break
	case info.PushType == "":
		return errors.New("推送类型不能为空")
		break
	}

	return nil
}

type Pusher interface {
	//验证推送数据是否符合要求
	Validate(data *Info) error
	//推送
	Push() (PushResponse, error)
}

func NewPusher(pusher_type string) (Pusher, error) {
	producer, ok := PusherRegisterList[pusher_type]
	if !ok {
		return nil, PUSHER_NOT_EXIST
	}

	return producer, nil
}
