package redis

import (
	"context"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/kit/redistools"
	"github.com/go-redis/redis/v8"
)

var masterRedisClient redistools.RedisClient

func InitRedis() error {
	var err error
	rc := redistools.RedisClient{
		Conf: conf.ConfRead{}.GetMasterRedisConf(),
	}
	err = rc.ConnRedis()
	if nil != err {
		return err
	}
	masterRedisClient = rc
	return nil
}

func GetMasterRedisClient() (client *redis.Client) {
	return masterRedisClient.Client
}

func LockMasterRedis(context context.Context, key string, outTime int) (bool, error) {
	return masterRedisClient.Lock(context, key, outTime)
}

func UnLockMasterRedis(context context.Context, key string) error {
	return masterRedisClient.UnLock(context, key)
}
