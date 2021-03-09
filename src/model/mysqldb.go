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
