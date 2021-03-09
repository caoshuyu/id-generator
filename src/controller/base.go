package controller

import (
	"github.com/caoshuyu/id-generator/src/cache/redis"
	"github.com/caoshuyu/id-generator/src/model"
)

type Controller struct {
}

func InitDb() {
	//初始化MySQL数据库信息
	model.GetMasterMysqlDb()
	//初始化Redis数据库信息
	err := redis.InitRedis()
	if nil != err {
		panic(err)
	}
}
