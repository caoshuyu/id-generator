package controller

import (
	"errors"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/id-generator/src/model"
	"strings"
)

func UpdateConf(ak, sk, name string) error {
	//校验配置信息
	var igcr conf.ConfRead
	var err error
	useAk, useSk := igcr.GetConfKey()
	if !strings.EqualFold(ak, useAk) || !strings.EqualFold(sk, useSk) {
		return errors.New("ak or sk error")
	}
	switch name {
	case "mysql.master":
		//检测新数据连接是否可用
		newConf, err := conf.ConfRead{}.NewConfGetMasterMysqlConf()
		if nil != err {
			return err
		}
		client, err := model.ConnectMysqlDb(newConf)
		if nil != err {
			return err
		}
		//更新数据库链接
		err = conf.UpdateConf(name)
		if nil != err {
			return err
		}
		model.UpdateMasterMysqlDb(client)

	case "log":
		err = conf.UpdateConf(name)
		if nil != err {
			return err
		}
	}

	return nil
}
