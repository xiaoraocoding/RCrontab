package main

import (
	"RCrontab/master"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.LoadHTMLGlob("web_root/*")
	//初始化配置文件
	master.InitConfig("./master.json")
	//初始化etcd的链接信息
	master.InitJobMgr()
	//初始化路由
	master.InitRouter(r)
	r.Run(":"+master.Conf.Port)



}