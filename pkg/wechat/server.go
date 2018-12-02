package wechat

import (
	"PushServer/pkg/setting"

	"sync"

	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
)

var wechats = make(map[string]*wechat.Wechat)

type MpInfo struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}

func (m *MpInfo) GetCustomService() *CustomService {
	return &CustomService{GetWechat(m)}
}

//获取一个wechat
//两重锁确保并发安全
func GetWechat(mp_info *MpInfo) *wechat.Wechat {
	wechat, ok := wechats[mp_info.AppID]

	if !ok {
		var mutex sync.Mutex
		mutex.Lock()
		defer mutex.Unlock()

		wechat, ok := wechats[mp_info.AppID]
		if !ok {
			wechat = createWechat(mp_info)
		}

		return wechat

	}

	return wechat
}

//创建一个wechat
func createWechat(mp_info *MpInfo) *wechat.Wechat {
	redisConfig := cache.RedisOpts{
		Host:        setting.CacheConfig.Host,
		Password:    setting.CacheConfig.Password,
		Database:    setting.CacheConfig.Database,
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 1,
	}

	redisCache := cache.NewRedis(&redisConfig)

	//配置微信参数
	config := &wechat.Config{
		AppID:          mp_info.AppID,
		AppSecret:      mp_info.AppSecret,
		Token:          mp_info.Token,
		EncodingAESKey: mp_info.EncodingAESKey,
		Cache:          redisCache,
	}

	wc := wechat.NewWechat(config)

	wechats[mp_info.AppID] = wc

	return wc
}
