package conf

const SERVER_NAME = "id-generator"

const (
	MYSQL_STATUS_USE    = 0
	MYSQL_STATUS_DELETE = 2
)

const (
	NULL_KEY_VALUE = "nil"
)

const (
	REDIS_KEY_PROJECT          = "idGenerator:"
	REDIS_KEY_PROJECT_INFO_KEY = REDIS_KEY_PROJECT + "projectKey:"
	REDIS_KEY_AUTO_NUMBER      = REDIS_KEY_PROJECT + "autoNumber:"
)

const (
	REDIS_LOCK_AUTO_NUMBER = "lock:" + REDIS_KEY_PROJECT + "autoNumber:"
)
