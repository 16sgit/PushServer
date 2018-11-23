package logging

import (
	"PushServer/pkg/setting"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitSetUp() {
	var js string

	if setting.LogConfig.Debug {
		js = fmt.Sprintf(`{
              "level": "%s",
              "encoding": "json",
              "outputPaths": ["stdout"],
              "errorOutputPaths": ["stdout"]
             }`, setting.LogConfig.LogLevel)
	} else {
		logPath := fmt.Sprintf("%s", setting.LogConfig.LogSavePath)
		if err := IsNotExistMkDir(logPath); err != nil {
			log.Fatal("open log file error: ", err)
		}

		logFileName := fmt.Sprintf("%s%s.%s", setting.LogConfig.LogSaveName, time.Now().Format(setting.LogConfig.TimeFormat),
			setting.LogConfig.LogFileExt)

		logOutPutPath := logPath + logFileName

		js = fmt.Sprintf(`{
              "level": "%s",
              "encoding": "json",
              "outputPaths": ["%s"],
              "errorOutputPaths": ["%s"]
            }`, setting.LogConfig.LogLevel, logOutPutPath, logOutPutPath)
	}

	log.Println(js)

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		log.Println(js)
		panic(js)
	}

	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error

	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}
