package controller

import "github.com/caoshuyu/id-generator/src/model"

type Controller struct {
}

func InitDb() {
	//初始化MySQL数据库信息
	model.GetMasterDb()
	//初始化Redis数据库信息

}
