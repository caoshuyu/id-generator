package model

import (
	"database/sql"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/kit/mysqltools"
)

var masterMysqlClient *mysqltools.MysqlClient

//获取链接
func GetMasterMysqlDb() *sql.DB {
	if nil == masterMysqlClient {
		client, err := ConnectMysqlDb(conf.ConfRead{}.GetMasterMysqlConf())
		if nil != err {
			panic(err)
		}
		masterMysqlClient = client
	}
	return masterMysqlClient.Client
}

func UpdateMasterMysqlDb(newClient *mysqltools.MysqlClient) {
	masterMysqlClient = newClient
}

