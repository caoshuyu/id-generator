package structure

type SetAutoNumber struct {
	IdType     string //id类型,int,string
	Filling    string //填充值，string类型可用，只一位
	ProjectId  string //项目标记
	TableName  string //表名称
	ColumnName string //列名称
	StPrefix   string //前缀
	NLength    int64  //Id长度
	StStart    int64  //开始id
	NIncrement int64  //步长
}

type GetAutoNumber struct {
	KeyNumber  string //主键ID
	ProjectId  string //项目标记
	TableName  string //表名称
	ColumnName string //列名称
}

type GetKeyNumber struct {
	ProjectId  string //项目标记
	TableName  string //表名称
	ColumnName string //列名称
}