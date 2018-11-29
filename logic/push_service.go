package logic

import (
	"PushServer/logic/push"
	"PushServer/models"
	"PushServer/pkg/cache"
	"fmt"

	//"PushServer/pkg/logging"
	"encoding/json"
	"log"
)

//将日志加入数据库
//满100条就将其加入数据库
func saveLogToDb(log_chan *chan models.TemplateMessageLog) {
	var i int = 0
	var logs []models.TemplateMessageLog
	for log_info := range *log_chan {
		i = i + 1
		logs = append(logs, log_info)

		if i%100 == 0 {
			models.AddTemplateMessageLogs(&logs)

			//清空数据
			i = 0
			logs = append([]models.TemplateMessageLog{})
		}
	}
}

//推送服务
func push_service() {
	var log_chan = make(chan models.TemplateMessageLog)
	go saveLogToDb(&log_chan)

	for {
		//获取推送数据
		push_info_string, err := cache.Pop()
		if err != nil {
			log.Fatal(err)
		}

		var info push.Info

		//验证数据格式是否符合要求
		json.Unmarshal([]byte(push_info_string), &info)
		if err = push.ValidateFormat(&info); err != nil {
			log.Fatal(err)
		}

		go func(info push.Info, log_chan *chan models.TemplateMessageLog) {
			//根据推送类型获取推送服务
			push_server, err := push.NewPusher(info.PushType)
			if err != nil {
				push_log := models.NewTemplateMessageLog(info.TemplateMessageId, "0", info.User, -1, fmt.Sprintf("%s", err))
				*log_chan <- push_log
				return
			}

			//验证是否符合推送要求
			if err = push_server.Validate(info); err != nil {
				push_log := models.NewTemplateMessageLog(info.TemplateMessageId, "0", info.User, -1, fmt.Sprintf("%s", err))
				*log_chan <- push_log
				return
			}

			//推送服务
			push_response, err := push_server.Push()
			if err != nil {
				log.Fatal(err)
			}

			push_log := models.NewTemplateMessageLog(info.TemplateMessageId, push_response.Msgid, info.User, push_response.Status, push_response.Message)
			*log_chan <- push_log

		}(info, &log_chan)
	}
}
