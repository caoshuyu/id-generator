package conf

var errorMap = map[string]int64{
	ProjectIdIsNull: 1000,
	IdTypeError:     1001,
	PTCIsHave:       1002,
}

const (
	ProjectIdIsNull = "project id is null"
	IdTypeError     = "id type not in int or string"
	PTCIsHave       = "project_id and table_name and column_name is have"
)

func GetErrorCode(err error) int64 {
	if nil == err {
		return 0
	}
	code, h := errorMap[err.Error()]
	if !h {
		return 2
	}
	return code
}
