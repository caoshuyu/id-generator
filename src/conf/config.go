package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/caoshuyu/kit/dlog"
	"github.com/caoshuyu/kit/mysqltools"
	"strconv"
	"time"
)

var requestHttpPort string
var sfc snowflakeConf
var mysql *mysqltools.MySqlConf
var confKey confKeyConf

func InitConf() {
	conf, err := getConf()
	if nil != err {
		panic(err)
	}
	requestHttpPort = ":" + strconv.Itoa(conf.Http.Port)
	initLog(&conf)
	mysql = initMysql(&conf)
	initConfKey(&conf)
}

func getConf() (config, error) {
	conf := config{}
	_, err := toml.DecodeFile("./"+SERVER_NAME+".toml", &conf)
	if nil != err {
		return conf, err
	}
	return conf, nil
}

func initLog(conf *config) {
	dlog.SetLog(dlog.SetLogConf{
		LogType: dlog.LOG_TYPE_LOCAL,
		LogPath: conf.Log.SavePath,
		Prefix:  SERVER_NAME,
	})
}

func initMysql(conf *config) *mysqltools.MySqlConf {
	return &mysqltools.MySqlConf{
		DbDsn:       conf.Mysql.Username + ":" + conf.Mysql.Password + "@tcp(" + conf.Mysql.Address + ")/" + conf.Mysql.DbName + "?" + conf.Mysql.Params,
		MaxOpen:     conf.Mysql.MaxOpen,
		MaxIdle:     conf.Mysql.MaxIdle,
		DbName:      conf.Mysql.DbName,
		MaxLifetime: time.Duration(conf.Mysql.MaxLifetime) * time.Second,
	}
}

func initConfKey(conf *config) {
	confKey = confKeyConf{
		Ak: conf.ConfKey.Ak,
		Sk: conf.ConfKey.Sk,
	}
}

//更新某个特定配置
func UpdateConf(confName string) (err error) {
	switch confName {
	case "mysql", "log":
	default:
		err = errors.New("conf name not have")
		return
	}
	conf, err := getConf()
	if nil != err {
		return err
	}
	switch confName {
	case "mysql":
		mysql = initMysql(&conf)
	case "log":
		initLog(&conf)
	}

	return
}

type ConfRead struct {
}

//新配置未生效MySQL配置
func (ConfRead) NewConfGetMysqlConf() (*mysqltools.MySqlConf, error) {
	conf, err := getConf()
	if nil != err {
		return nil, err
	}
	return initMysql(&conf), nil
}

func (ConfRead) GetRequestHttpPort() string {
	return requestHttpPort
}

func (ConfRead) GetSnowflakeConf() (dataCenterId, machineId int64) {
	return sfc.DataCenterId, sfc.MachineId
}

func (ConfRead) GetMysqlConf() *mysqltools.MySqlConf {
	return mysql
}

func (ConfRead) GetConfKey() (ak, sk string) {
	return confKey.Ak, confKey.Sk
}
