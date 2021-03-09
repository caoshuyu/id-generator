package model

import (
	"database/sql"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/kit/mysqltools"
)

var masterDbClient *mysqltools.MysqlClient

//获取链接
func GetMasterDb() *sql.DB {
	if nil == masterDbClient {
		client, err := ConnectMysqlDb(conf.ConfRead{}.GetMysqlConf())
		if nil != err {
			panic(err)
		}
		masterDbClient = client
	}
	return masterDbClient.Client
}

func UpdateMasterDb(newClient *mysqltools.MysqlClient) {
	masterDbClient = newClient
}

func UpdateTableValue(db *sql.DB, sqlText string, params []interface{}) (number int64, err error) {
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		params...,
	)
	number, err = res.RowsAffected()
	if nil != err {
		return 0, err
	}
	return number, nil
}
