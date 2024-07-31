package core

import (
	"github.com/Linxhhh/easy-doc/global"
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func InitRedis(db int) *redis.Client {

	redisConf := global.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		PoolSize: redisConf.PoolSize,
		DB:       db,
	})

	// 设置超时时间
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	
	// 测试 Redis 连接是否正常
	_, err := client.Ping().Result()
	if err != nil {
		logrus.Fatalf("%s Redis 连接失败，%s", redisConf.Addr(), err.Error())
	}

	return client
}