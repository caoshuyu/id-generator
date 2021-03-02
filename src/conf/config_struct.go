package conf

type config struct {
	Mysql     mysqlConf
	Http      httpConf
	Log       logConf
	Snowflake snowflakeConf
	ConfKey   confKeyConf
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
