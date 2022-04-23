package main

import (
	"RCrontab/master"
	"github.com/gin-gonic/gin"
)

func main() {
	//开启web服务
	r := gin.Default()
    //渲染模版
	r.LoadHTMLGlob("web_root/*")
	//初始化配置文件
	master.InitConfig("./master.json")
	//初始化etcd的链接信息
	master.InitJobMgr()
    //初始化路由
	master.InitRouter(r)
	master.InitLogMgr()
	r.Run(":"+master.Conf.Port)



}
