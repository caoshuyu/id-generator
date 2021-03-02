package request_http

import (
	"github.com/caoshuyu/id-generator/src/controller"
	"github.com/labstack/echo/v4"
)

type businessRouterFunc struct {
}

var brf businessRouterFunc

var contStruct controller.Controller

//获取通用ID，使用雪花算法（snowflake）
func (*businessRouterFunc) getIdUseSnowflake(ectx echo.Context) (err error) {



	return nil
}

//设置项目自增ID信息
func (*businessRouterFunc) setAutoId(ectx echo.Context) (err error) {

	return nil
}

//根据project_id，table_name，column_name获取标记Key
func (*businessRouterFunc) getAutoIdKey(ectx echo.Context) (err error) {

	return nil
}

//获取自增ID
func (*businessRouterFunc) getAutoId(ectx echo.Context) (err error) {

	return nil
}



