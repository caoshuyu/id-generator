package request_http

import (
	"fmt"
	"github.com/caoshuyu/id-generator/src/conf"
	"github.com/caoshuyu/kit/echomiddleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func ListeningHTTP() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders:     []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	e.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "ping")
	})

	router(e)

	err := e.StartServer(&http.Server{
		Addr:              conf.ConfRead{}.GetRequestHttpPort(),
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 2,
		WriteTimeout:      time.Second * 30,
	})
	if nil != err {
		fmt.Println("server_start_err", err)
	}

}

func router(e *echo.Echo) {
	e.Use(echomiddleware.Gls, echomiddleware.Access, echomiddleware.Recover)
	confRouter(e)
	businessRouter(e)
}

//业务接口
func businessRouter(e *echo.Echo) {
	//获取通用ID，使用雪花算法（snowflake）
	e.Any("/snowflake_id", brf.getIdUseSnowflake)
	//设置项目自增ID信息
	e.POST("/set_auto_id", brf.setAutoId)
	//根据project_id，table_name，column_name获取标记Key
	e.POST("/get_auto_id_key", brf.getAutoIdKey)
	//获取自增ID
	e.POST("/get_auto_number", brf.getAutoNumber)
}

//配置接口
func confRouter(e *echo.Echo) {
	g := e.Group("/conf")
	g.GET("/update_conf", crf.updateConf)
}
