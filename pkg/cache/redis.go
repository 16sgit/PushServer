package cache

import (
	"PushServer/pkg/setting"
	"errors"
	"log"
	"time"

	"github.com/Unknwon/com"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func InitSetUp() {
	client = redis.NewClient(&redis.Options{
		Addr:     setting.CacheConfig.Host,
		Password: setting.CacheConfig.Password,
		DB:       setting.CacheConfig.Database,
	})

	ping, err := client.Ping().Result()

	if err != nil {
		log.Fatal(ping, err)
	}
}

//从队列中获取一条数据
func Pop() (string, error) {
	val, err := client.BRPop(setting.CacheConfig.Timeout, setting.CacheConfig.QueueName).Result()

	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return val[1], nil
}

//添加数据
func Push(s string) error {
	return client.LPush(setting.CacheConfig.QueueName, s).Err()
}

//递增hash的某个字段的值，并返回当前缓存的值
//为其设置过期时间
func IncreaseHash(key, field string, offset int64, expire time.Duration) (int64, error) {
	var is_first bool = false
	if client.Exists(key).Val() == 0 {
		is_first = true
	}

	num, err := client.HIncrBy(key, field, offset).Result()
	if err != nil {
		return num, err
	}

	if is_first {
		if !client.Expire(key, expire).Val() {
			return num, errors.New("cache:设置过期时间失败")
		}
	}

	return num, nil
}

//当前分钟内的刷卡数量值+1
//缓存保存两天
func IncreaseConsumptionNumOfPerMinute() error {
	now := time.Now()
	now_date := now.Format("2006-01-02")
	now_minute_timestamp := now.Unix() - int64(now.Second())
	key := "per_minute:" + com.ToStr(now_date)
	field := com.ToStr(now_minute_timestamp)
	_, err := IncreaseHash(key, field, 1, time.Hour*24*2)
	return err
}
