package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/caoshuyu/kit/dlog"
	"github.com/caoshuyu/kit/mysqltools"
	"github.com/caoshuyu/kit/redistools"
	"strconv"
	"time"
)

var requestHttpPort string
var sfc snowflakeConf
var masterMysql *mysqltools.MySqlConf
var masterRedis *redistools.RedisConf
var confKey confKeyConf
var throughAttackTime int64

func InitConf() {
	conf, err := getConf()
	if nil != err {
		panic(err)
	}
	requestHttpPort = ":" + strconv.Itoa(conf.Http.Port)
	initLog(&conf)
	masterMysql = initMasterMysql(&conf)
	masterRedis = initMasterRedis(&conf)
	initConfKey(&conf)
	initThroughAttack(&conf)
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

func initMasterMysql(conf *config) *mysqltools.MySqlConf {
	return &mysqltools.MySqlConf{
		DbDsn: conf.Mysql.Master.Username + ":" + conf.Mysql.Master.Password + "@tcp(" + conf.Mysql.Master.Address + ")/" +
			conf.Mysql.Master.DbName + "?" + conf.Mysql.Master.Params,
		MaxOpen:     conf.Mysql.Master.MaxOpen,
		MaxIdle:     conf.Mysql.Master.MaxIdle,
		DbName:      conf.Mysql.Master.DbName,
		MaxLifetime: time.Duration(conf.Mysql.Master.MaxLifetime) * time.Second,
	}
}

func initMasterRedis(conf *config) *redistools.RedisConf {
	return &redistools.RedisConf{
		Addr:     conf.Redis.Master.Addr,
		Password: conf.Redis.Master.Password,
		DB:       conf.Redis.Master.DB,
	}
}

func initConfKey(conf *config) {
	confKey = confKeyConf{
		Ak: conf.ConfKey.Ak,
		Sk: conf.ConfKey.Sk,
	}
}

func initThroughAttack(conf *config) {
	throughAttackTime = conf.ThroughAttack.TimeSecond
}

//更新某个特定配置
func UpdateConf(confName string) (err error) {
	switch confName {
	case "mysql.master", "log":
	default:
		err = errors.New("conf name not have")
		return
	}
	conf, err := getConf()
	if nil != err {
		return err
	}
	switch confName {
	case "mysql.master":
		masterMysql = initMasterMysql(&conf)
	case "log":
		initLog(&conf)
	case "through_attack":
		initThroughAttack(&conf)
	}

	return
}

type ConfRead struct {
}

//新配置未生效MySQL配置
func (ConfRead) NewConfGetMasterMysqlConf() (*mysqltools.MySqlConf, error) {
	conf, err := getConf()
	if nil != err {
		return nil, err
	}
	return initMasterMysql(&conf), nil
}

func (ConfRead) GetRequestHttpPort() string {
	return requestHttpPort
}

func (ConfRead) GetSnowflakeConf() (dataCenterId, machineId int64) {
	return sfc.DataCenterId, sfc.MachineId
}

func (ConfRead) GetMasterMysqlConf() *mysqltools.MySqlConf {
	return masterMysql
}

func (ConfRead) GetMasterRedisConf() *redistools.RedisConf {
	return masterRedis
}

func (ConfRead) GetConfKey() (ak, sk string) {
	return confKey.Ak, confKey.Sk
}

func (ConfRead) GetThroughAttackTime() int64 {
	return throughAttackTime
}
