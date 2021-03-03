package controller

import (
	"context"
	"errors"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/id-generator/src/controller/structure"
	"github.com/caoshuyu/id-generator/src/model"
	"github.com/caoshuyu/kit/global_const"
	"github.com/caoshuyu/kit/stringtools"
	"strconv"
	"strings"
)

func (c *Controller) SetAutoNumber(ectx context.Context, input *structure.SetAutoNumber) (out string, err error) {
	//check params
	if strings.EqualFold(input.ProjectId, "") {
		err = errors.New(conf.ProjectIdIsNull)
		return
	}
	switch input.IdType {
	case "int":
		input.Filling = ""
		input.StPrefix = ""
	case "string":
		switch len(input.Filling) {
		case 0:
			input.Filling = "0"
		case 1:
		default:
			input.Filling = input.Filling[:1]
		}
	default:
		err = errors.New(conf.IdTypeError)
		return
	}
	if input.StStart == 0 {
		input.StStart = 1
	}
	if input.NIncrement == 0 {
		input.NIncrement = 500
	}

	//check projectId,tableName,columnName is have
	autoIdModel := model.AutoId{}
	db := model.GetMasterDb()
	number, err := autoIdModel.CountByProjectIdAndTableNameAndColumnName(db, input.ProjectId, input.TableName, input.ColumnName)
	if nil != err {
		return
	}
	if number > 0 {
		err = errors.New(conf.PTCIsHave)
		return
	}
	//get id
	id, err := c.GetIdUseSnowflake()
	if nil != err {
		return
	}
	//save table
	autoId := model.AutoId{
		Id:         id,
		ProjectId:  input.ProjectId,
		TableName:  input.TableName,
		ColumnName: input.ColumnName,
		IdType:     input.IdType,
		Filling:    input.Filling,
		StPrefix:   input.StPrefix,
		NLength:    input.NLength,
		StStart:    input.StStart,
		StNow:      input.StStart,
		NIncrement: input.NIncrement,
		State:      global_const.MYSQL_STATUS_USE,
	}
	err = autoId.Insert(db)
	if nil != err {
		return
	}
	//return value
	out = strconv.FormatInt(id, 10)
	return
}

func (c *Controller) GetAutoNumber(ectx context.Context, input *structure.GetAutoNumber) (out string, err error) {
	var useKey bool
	if len(input.KeyNumber) > 0 {
		useKey = true
	}
	if !useKey {
		if strings.EqualFold(input.ProjectId, "") {
			err = errors.New(conf.ProjectIdIsNull)
			return
		}
	}
	var redisSaveKey string
	if useKey {
		redisSaveKey = "autoNumber:" + stringtools.Md5(input.KeyNumber)
	} else {
		redisSaveKey = "autoNumber:" + stringtools.Md5(input.ProjectId+"_"+input.TableName+"_"+input.ColumnName)
	}

	_ = redisSaveKey
	


	return
}

func (c *Controller) GetKeyNumber(ectx context.Context, input *structure.GetKeyNumber) (out string, err error) {
	if strings.EqualFold(input.ProjectId, "") {
		err = errors.New(conf.ProjectIdIsNull)
		return
	}
	autoIdModel := model.AutoId{}
	db := model.GetMasterDb()
	id, err := autoIdModel.SelectIdByProjectIdAndTableNameAndColumnName(db, input.ProjectId, input.TableName, input.ColumnName)
	if nil != err {
		return
	}
	return strconv.FormatInt(id, 10), nil
}
