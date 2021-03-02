package controller

import (
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/kit/dlog"
	"github.com/caoshuyu/snowflake"
)

var snowFlack snowflake.SnowFlack

func init() {
	dataCenterId, machineId := conf.ConfRead{}.GetSnowflakeConf()
	snowFlack = snowflake.SnowFlack{
		DatacenterId: dataCenterId,
		MachineId:    machineId,
	}
	err := snowFlack.Init()
	if nil != err {
		panic(err)
		return
	}
}

func (*Controller) GetIdUseSnowflake() (id int64, err error) {
	id, err = snowFlack.NextId()
	if nil != err {
		dlog.ERROR("funcName", "GetIdUseSnowflake", "error", err)
		return
	}
	return
}
