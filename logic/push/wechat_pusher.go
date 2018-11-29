package push

type WeChatPusher struct {
}

//验证推送数据是否符合要求
func (p *WeChatPusher) Validate(data Info) error {
	return nil
}

//推送
func (p WeChatPusher) Push() (PushResponse, error) {
	return PushResponse{}, nil
}
