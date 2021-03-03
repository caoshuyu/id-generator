package request_http

type GetIdUseSnowflakeResponse struct {
	Number int64 `json:"number"` //生成的标识
}

type SetAutoIdRequest struct {
	IdType     string `json:"id_type"`     //id类型,int,string
	Filling    string `json:"filling"`     //填充值，string类型可用，只一位
	ProjectId  string `json:"project_id"`  //项目标记
	TableName  string `json:"table_name"`  //表名称
	ColumnName string `json:"column_name"` //列名称
	StPrefix   string `json:"st_prefix"`   //前缀
	NLength    int64  `json:"n_length"`    //Id长度
	StStart    int64  `json:"st_start"`    //开始id
	NIncrement int64  `json:"n_increment"` //步长
}

type SetAutoIdResponse struct {
	KeyNumber string `json:"key_number"`
}

type GetAutoIdKeyRequest struct {
	ProjectId  string `json:"project_id"`  //项目标记
	TableName  string `json:"table_name"`  //表名称
	ColumnName string `json:"column_name"` //列名称
}

type GetAutoIdKeyResponse struct {
	KeyNumber string `json:"key_number"`
}

type GetAutoNumberRequest struct {
	KeyNumber  string `json:"key_number"`
	ProjectId  string `json:"project_id"`  //项目标记
	TableName  string `json:"table_name"`  //表名称
	ColumnName string `json:"column_name"` //列名称
}

type GetAutoNumberResponse struct {
	Number string `json:"number"` //生成的标识
}
