package models

import (
	"PushServer/pkg/setting"
	"strings"
	"time"

	"github.com/Unknwon/com"
)

type TemplateMessageLog struct {
	ID      int64 `gorm:"primary_key" json:"id"`
	Tid     int64
	Msgid   string
	User    string
	Status  int
	Message string
	Time    int64
}

func NewTemplateMessageLog(tid int64, msgid string, user string, status int, message string) TemplateMessageLog {
	return TemplateMessageLog{
		Tid:     tid,
		Msgid:   msgid,
		User:    user,
		Status:  status,
		Message: message,
		Time:    time.Now().Unix(),
	}
}

//批量加入数据库
func AddTemplateMessageLogs(logs *[]TemplateMessageLog) error {
	var teble string = setting.DbConfig.TablePrefix + "template_message_log"
	//使用原生sql添加数据
	var sql string = "INSERT INTO `" + teble + "` (`tid`, `msgid`,`user`, `status`, `message`,`time`) VALUES"
	for _, log := range *logs {
		sql += "(" + com.ToStr(log.Tid, 10) + ",'"
		sql += com.ToStr(log.Msgid, 10) + "','"
		sql += log.User + "',"
		sql += com.ToStr(log.Status, 10) + ",'"
		sql += log.Message + "',"
		sql += com.ToStr(log.Time) + "),"
	}
	sql = strings.TrimRight(sql, ",")
	db.Exec(sql)

	return nil
}
