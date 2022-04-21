package main

import (
	"RCrontab/master"
	"github.com/gin-gonic/gin"
)

func main() {
	//开启web服务
		r := gin.Default()

	     r.LoadHTMLGlob("web_root/*")
		master.InitConfig("./master.json")
	    master.InitJobMgr()

		master.InitRouter(r)
		r.Run(":"+master.Conf.Port)



}
