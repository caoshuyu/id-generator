package model

import (
	"database/sql"
	"github.com/caoshuyu/kit/mysqltools"
	"time"
)

type AutoId struct {
	Id         int64
	ProjectId  string
	TableName  string
	ColumnName string
	IdType     string
	Filling    string
	StPrefix   string
	NLength    int64
	StStart    int64
	StNow      int64
	NIncrement int64
	State      int64
	CreateAt   time.Time
	UpdateAt   time.Time
}

func (t *AutoId) tableName() string {
	return "auto_id"
}

func (t *AutoId) Insert(db *sql.DB) error {
	sqlText := "INSERT INTO " + t.tableName() +
		" (id,project_id,table_name,column_name,id_type,filling,st_prefix,n_length,st_start,st_now,n_increment,state) " +
		" VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		mysqltools.WriteDbError(t.tableName(), err, sqlText)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		t.Id,
		t.ProjectId,
		t.TableName,
		t.ColumnName,
		t.IdType,
		t.Filling,
		t.StPrefix,
		t.NLength,
		t.StStart,
		t.StNow,
		t.NIncrement,
		t.State,
	)
	if nil != err {
		t._errorInsert(err, sqlText)
		return err
	}
	return nil
}

func (t *AutoId) _errorInsert(err error, sqlText string) {
	errMap := make(map[string]interface{})
	errMap["id"] = t.Id
	errMap["project_id"] = t.ProjectId
	errMap["table_name"] = t.TableName
	errMap["column_name"] = t.ColumnName
	errMap["id_type"] = t.IdType
	errMap["filling"] = t.Filling
	errMap["st_prefix"] = t.StPrefix
	errMap["n_length"] = t.NLength
	errMap["st_start"] = t.StStart
	errMap["st_now"] = t.StNow
	errMap["n_increment"] = t.NIncrement
	errMap["state"] = t.State
	mysqltools.WriteDbError(t.tableName(), err, sqlText, errMap)
}

func (t *AutoId) CountByProjectIdAndTableNameAndColumnName(db *sql.DB, projectId string, tableName string, columnName string) (int64, error) {
	var count int64
	sqlText := "SELECT COUNT(*) as count FROM " + t.tableName() + " WHERE project_id = ? AND table_name = ? AND column_name = ?"
	err := db.QueryRow(sqlText,
		projectId,
		tableName,
		columnName).Scan(
		&count,
	)
	if nil != err {
		t._errorCountByProjectIdAndTableNameAndColumnName(err, sqlText, projectId, tableName, columnName)
		return 0, err
	}
	return count, nil
}

func (t *AutoId) _errorCountByProjectIdAndTableNameAndColumnName(err error, sqlText string, projectId string, tableName string, columnName string) {
	errMap := make(map[string]interface{})
	errMap["project_id"] = projectId
	errMap["table_name"] = tableName
	errMap["column_name"] = columnName
	mysqltools.WriteDbError(t.tableName(), err, sqlText, errMap)
}

func (t *AutoId) SelectIdByProjectIdAndTableNameAndColumnName(db *sql.DB, projectId string, tableName string, columnName string) (int64, error) {
	var id int64
	sqlText := "SELECT id FROM " + t.tableName() + " WHERE project_id = ? AND table_name = ? AND column_name = ?"
	err := db.QueryRow(sqlText,
		projectId,
		tableName,
		columnName).Scan(
		&id,
	)
	if nil != err {
		t._errorSelectIdByProjectIdAndTableNameAndColumnName(err, sqlText, projectId, tableName, columnName)
		return 0, err
	}
	return id, nil
}

func (t *AutoId) _errorSelectIdByProjectIdAndTableNameAndColumnName(err error, sqlText string, projectId string, tableName string, columnName string) {
	errMap := make(map[string]interface{})
	errMap["project_id"] = projectId
	errMap["table_name"] = tableName
	errMap["column_name"] = columnName
	mysqltools.WriteDbError(t.tableName(), err, sqlText, errMap)
}

func (t *AutoId) SelectById(db *sql.DB, id int64) (*AutoId, error) {
	sqlText := "SELECT id_type,filling,st_prefix,n_length,st_now,n_increment FROM " + t.tableName() + " WHERE id = ?"
	ai := AutoId{}
	err := db.QueryRow(sqlText, id).Scan(
		&ai.IdType,
		&ai.Filling,
		&ai.StPrefix,
		&ai.NLength,
		&ai.StNow,
		&ai.NIncrement,
	)
	if nil != err {
		t._errorSelectById(err, sqlText, id)
		return nil, err
	}
	return &ai, nil
}
func (t *AutoId) _errorSelectById(err error, sqlText string, id int64) {
	errMap := make(map[string]interface{})
	errMap["id"] = id
	mysqltools.WriteDbError(t.tableName(), err, sqlText, errMap)
}

func (t *AutoId) UpdateStNowById(db *sql.DB, id int64) (number int64, err error) {
	sqlText := "UPDATE auto_id SET st_now = st_now + n_increment WHERE id = ?"
	var params []interface{}
	params = append(params, id)
	number, err = UpdateTableValue(db, sqlText, params)
	if nil != err {
		t._errorUpdateStNowById(err, sqlText, id)
		return 0, err
	}
	return number, nil
}
func (t *AutoId) _errorUpdateStNowById(err error, sqlText string, id int64) {
	errMap := make(map[string]interface{})
	errMap["id"] = id
	mysqltools.WriteDbError(t.tableName(), err, sqlText, errMap)
}
