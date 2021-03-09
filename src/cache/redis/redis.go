package redis

import (
	"context"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/kit/redistools"
	"github.com/go-redis/redis/v8"
)

var redisClient redistools.RedisClient

func InitRedis() error {
	rc := redistools.RedisClient{
		Conf: conf.ConfRead{}.GetRedisConf(),
	}
	err := rc.ConnRedis()
	if nil != err {
		return err
	}
	redisClient = rc
	return nil
}

func GetRedisClient() (client *redis.Client) {
	return redisClient.Client
}

func LockRedis(context context.Context, key string, outTime int) (bool, error) {
	return redisClient.Lock(context, key, outTime)
}

func UnLockRedis(context context.Context, key string) error {
	return redisClient.UnLock(context, key)
}
