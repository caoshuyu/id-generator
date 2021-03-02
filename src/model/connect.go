//链接数据库
package model

import "github.com/caoshuyu/kit/mysqltools"

func ConnectMysqlDb(conf *mysqltools.MySqlConf) (client *mysqltools.MysqlClient, err error) {
	client = &mysqltools.MysqlClient{
		Conf: conf,
	}
	err = client.Connect()
	if nil != err {
		return
	}
	return
}
