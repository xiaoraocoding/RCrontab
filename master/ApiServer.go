package master

import (
	"RCrontab/common"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)


func InitRouter(r *gin.Engine)  {
	r.POST("/job/save",HanleJobSave)

}


//保存任务端口
func HanleJobSave (c *gin.Context) {
    var job common.Job
	postJob := c.PostForm("job")
	err := json.Unmarshal([]byte(postJob),&job)
	if err != nil {
		fmt.Println("json unmarshal failed err:",err)
	}
	oldJob,err := G_JobMgr.SaveJob(&job)
	if err != nil {
		fmt.Println("save etcd failed,err:",err)
	}
	c.JSON(200,gin.H{
		"error": 0,
		"msg" : "success",
		"oldJob":oldJob,
	})








}

