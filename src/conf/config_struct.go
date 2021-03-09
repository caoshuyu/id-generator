package conf

type config struct {
	Mysql         mysqlConf
	Redis         redisConf
	Http          httpConf
	Log           logConf
	Snowflake     snowflakeConf
	ConfKey       confKeyConf       `toml:"conf_key"`
	ThroughAttack throughAttackConf `toml:"through_attack"`
}

type mysqlConf struct {
	Username    string
	Password    string
	Address     string
	DbName      string `toml:"db_name"`
	Params      string
	MaxOpen     int `toml:"max_open"`
	MaxIdle     int `toml:"max_idle"`
	MaxLifetime int `toml:"max_lifetime"`
}

type httpConf struct {
	Port int
}

type logConf struct {
	SavePath string `toml:"save_path"`
}

type snowflakeConf struct {
	DataCenterId int64 `toml:"data_center_id"`
	MachineId    int64 `toml:"machine_id"`
}

type confKeyConf struct {
	Ak string
	Sk string
}

type redisConf struct {
	Addr     string
	Password string
	DB       int
}

type throughAttackConf struct {
	TimeSecond int64 `toml:"time_second"`
}
