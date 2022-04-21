package main

import (
	"RCrontab/master"
	"github.com/gin-gonic/gin"
)

func main() {
	//开启web服务
		r := gin.Default()
		master.InitConfig("./master.json")

		master.InitRouter(r)
		r.Run(":"+master.Conf.Port)



}
