package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

//获取错误消息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]

	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
