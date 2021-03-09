package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/caoshuyu/id-generator/src/cache/redis"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/id-generator/src/controller/structure"
	"github.com/caoshuyu/id-generator/src/model"
	"github.com/caoshuyu/kit/global_const"
	"github.com/caoshuyu/kit/stringtools"
	"strconv"
	"strings"
	"time"
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
	if input.StStart <= 0 {
		input.StStart = 1
	}
	if input.NIncrement == 0 {
		input.NIncrement = 500
	}

	//check projectId,tableName,columnName is have
	autoIdModel := model.AutoId{}
	db := model.GetMasterMysqlDb()
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
		StNow:      input.StStart - 1,
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
		_, err = strconv.ParseInt(input.KeyNumber, 10, 64)
		if nil != err {
			return "", errors.New(conf.ProjectIdNotHave)
		}
		useKey = true
	}
	if !useKey {
		if strings.EqualFold(input.ProjectId, "") {
			return "", errors.New(conf.ProjectIdIsNull)
		}
		//project info to key
		input.KeyNumber, err = _projectKey(ectx, input.ProjectId, input.TableName, input.ColumnName)
		if nil != err {
			return "", err
		}
	}
	if strings.EqualFold(input.KeyNumber, conf.NULL_KEY_VALUE) {
		return "", errors.New(conf.ProjectIdNotHave)
	}

	value := _getValueNumber(ectx, input.KeyNumber)
	if !strings.EqualFold(value, "") {
		return value, nil
	}
	// not have key value get it
	err = _initValueNumber(ectx, input.KeyNumber)
	if nil != err {
		return "", err
	}
	return _getValueNumber(ectx, input.KeyNumber), nil
}

func _getValueNumber(ectx context.Context, key string) string {
	redisUseKey := conf.REDIS_KEY_AUTO_NUMBER + stringtools.Md5(key)
	valNumber := redis.GetMasterRedisClient().LPop(ectx, redisUseKey).Val()
	if len(valNumber) > 0 {
		return valNumber
	}
	return ""
}

func _initValueNumber(ectx context.Context, key string) error {
	redisLockKey := conf.REDIS_LOCK_AUTO_NUMBER + key
	b, err := redis.LockMasterRedis(ectx, redisLockKey, 6)
	if nil != err {
		return err
	}
	if !b {
		return errors.New(conf.KeyLocked)
	}
	defer redis.UnLockMasterRedis(ectx, redisLockKey)
	ai := model.AutoId{}
	db := model.GetMasterMysqlDb()

	id, err := strconv.ParseInt(key, 10, 64)
	if nil != err {
		return err
	}
	aiVal, err := ai.SelectById(db, id)
	if nil != err {
		if sql.ErrNoRows == err {
			return errors.New(conf.ProjectIdNotHave)
		}
		return err
	}
	n, err := ai.UpdateStNowById(db, id)
	if nil != err {
		return err
	}
	if n == 0 {
		return errors.New(conf.ProjectIdNotHave)
	}
	numberList := make([]interface{}, 0, aiVal.NIncrement)
	switch aiVal.IdType {
	case "int":
		for i, j := aiVal.StNow+1, aiVal.StNow+aiVal.NIncrement; i <= j; i++ {
			numberList = append(numberList, i)
		}
	case "string":
		l := aiVal.NLength - int64(len(aiVal.StPrefix))
		for i, j := aiVal.StNow+1, aiVal.StNow+aiVal.NIncrement; i <= j; i++ {
			var number string
			if strings.EqualFold("0", aiVal.Filling) {
				number = aiVal.StPrefix + fmt.Sprintf("%0"+strconv.FormatInt(l, 10)+"d", i)
			} else {
				numberList = append(numberList, aiVal.StPrefix+strings.Repeat(aiVal.Filling, int(l)-len(strconv.FormatInt(i, 10)))+strconv.FormatInt(i, 10))
			}
			numberList = append(numberList, number)
		}
	}

	redisUseKey := conf.REDIS_KEY_AUTO_NUMBER + stringtools.Md5(key)
	return redis.GetMasterRedisClient().RPush(ectx, redisUseKey, numberList...).Err()
}

func _projectKey(ectx context.Context, projectId, tableName, columnName string) (string, error) {
	redisKey := conf.REDIS_KEY_PROJECT_INFO_KEY + stringtools.Md5(projectId+"_"+tableName+"_"+columnName)
	val := redis.GetMasterRedisClient().Get(ectx, redisKey).Val()
	if !strings.EqualFold(val, "") {
		return val, nil
	}
	db := model.GetMasterMysqlDb()
	ai := model.AutoId{}
	id, err := ai.SelectIdByProjectIdAndTableNameAndColumnName(db, projectId, tableName, columnName)
	if nil != err {
		if err == sql.ErrNoRows {
			redis.GetMasterRedisClient().Set(ectx, redisKey, conf.NULL_KEY_VALUE, time.Duration(conf.ConfRead{}.GetThroughAttackTime())*time.Second)
			return "", errors.New(conf.ProjectIdIsNull)
		}
		return "", err
	}
	key := strconv.FormatInt(id, 10)
	redis.GetMasterRedisClient().Set(ectx, redisKey, key, time.Duration(60)*time.Minute)
	return key, nil
}

func (c *Controller) GetKeyNumber(ectx context.Context, input *structure.GetKeyNumber) (out string, err error) {
	if strings.EqualFold(input.ProjectId, "") {
		err = errors.New(conf.ProjectIdIsNull)
		return
	}
	autoIdModel := model.AutoId{}
	db := model.GetMasterMysqlDb()
	id, err := autoIdModel.SelectIdByProjectIdAndTableNameAndColumnName(db, input.ProjectId, input.TableName, input.ColumnName)
	if nil != err {
		return
	}
	return strconv.FormatInt(id, 10), nil
}
