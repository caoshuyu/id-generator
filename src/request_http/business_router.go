package request_http

import (
	"context"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/id-generator/src/controller"
	"github.com/caoshuyu/id-generator/src/controller/structure"
	"github.com/caoshuyu/kit/echo_out_tools"
	"github.com/labstack/echo/v4"
)

type businessRouterFunc struct {
}

var brf businessRouterFunc

var contStruct controller.Controller

//获取通用ID，使用雪花算法（snowflake）
func (*businessRouterFunc) getIdUseSnowflake(ectx echo.Context) (err error) {
	id, err := contStruct.GetIdUseSnowflake()
	if nil != err {
		return echo_out_tools.EchoErrorData(ectx, err, 2)
	}
	return echo_out_tools.EchoSuccessData(ectx, GetIdUseSnowflakeResponse{
		Number: id,
	})
}

//设置项目自增ID信息
func (*businessRouterFunc) setAutoId(ectx echo.Context) (err error) {
	req := new(SetAutoIdRequest)
	if err = ectx.Bind(req); err != nil {
		return echo_out_tools.EchoErrorData(ectx, err, 2)
	}
	input := structure.SetAutoNumber{
		IdType:     req.IdType,
		Filling:    req.Filling,
		ProjectId:  req.ProjectId,
		TableName:  req.TableName,
		ColumnName: req.ColumnName,
		StPrefix:   req.StPrefix,
		NLength:    req.NLength,
		StStart:    req.StStart,
		NIncrement: req.NIncrement,
	}
	var ctx context.Context
	id, err := contStruct.SetAutoNumber(ctx, &input)
	if nil != err {
		return echo_out_tools.EchoErrorData(ectx, err, conf.GetErrorCode(err))
	}
	return echo_out_tools.EchoSuccessData(ectx, SetAutoIdResponse{
		KeyNumber: id,
	})
}

//根据project_id，table_name，column_name获取标记Key
func (*businessRouterFunc) getAutoIdKey(ectx echo.Context) (err error) {
	req := new(GetAutoIdKeyRequest)
	if err = ectx.Bind(req); err != nil {
		return echo_out_tools.EchoErrorData(ectx, err, 2)
	}
	input := structure.GetKeyNumber{
		ProjectId:  req.ProjectId,
		TableName:  req.TableName,
		ColumnName: req.ColumnName,
	}
	var ctx context.Context
	id, err := contStruct.GetKeyNumber(ctx, &input)
	if nil != err {
		return echo_out_tools.EchoErrorData(ectx, err, conf.GetErrorCode(err))
	}
	return echo_out_tools.EchoSuccessData(ectx, GetAutoIdKeyResponse{
		KeyNumber: id,
	})
}

//获取自增ID
func (*businessRouterFunc) getAutoNumber(ectx echo.Context) (err error) {
	req := new(GetAutoNumberRequest)
	if err = ectx.Bind(req); err != nil {
		return echo_out_tools.EchoErrorData(ectx, err, 2)
	}
	input := structure.GetAutoNumber{
		KeyNumber:  req.KeyNumber,
		ProjectId:  req.ProjectId,
		TableName:  req.TableName,
		ColumnName: req.ColumnName,
	}
	var ctx context.Context
	number, err := contStruct.GetAutoNumber(ctx, &input)
	if nil != err {
		return echo_out_tools.EchoErrorData(ectx, err, conf.GetErrorCode(err))
	}
	return echo_out_tools.EchoSuccessData(ectx, GetAutoNumberResponse{
		Number: number,
	})
}
