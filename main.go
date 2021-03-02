package main

import (
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/id-generator/src/controller"
	"github.com/caoshuyu/id-generator/src/request_http"
)

func main() {
	//初始化配置信息
	conf.InitConf()
	//初始化数据库信息
	controller.InitDb()

	//启动HTTP服务
	request_http.ListeningHTTP()
}
