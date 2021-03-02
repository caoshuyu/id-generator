package request_http

import (
	"github.com/caoshuyu/id-generator/src/controller"
	"github.com/caoshuyu/kit/echo_out_tools"
	"github.com/labstack/echo/v4"
)

type confRouterFunc struct {
}

var crf confRouterFunc

func (*confRouterFunc) updateConf(ectx echo.Context) (err error) {
	//获取校验参数
	ak := ectx.Request().Header.Get("ak")
	sk := ectx.Request().Header.Get("sk")
	name := ectx.QueryParam("name")
	err = controller.UpdateConf(ak, sk, name)
	if nil != err {
		return echo_out_tools.EchoErrorData(ectx, err, 2)
	}
	return echo_out_tools.EchoSuccessData(ectx, "")
}
