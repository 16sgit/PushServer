package cache

import (
	"PushServer/pkg/setting"

	"log"

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
