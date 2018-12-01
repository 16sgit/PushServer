//注册者列表
package push

var PusherRegisterList = map[string]Pusher{
	"wechat_template": &WeChatTemplatePusher{},
}
