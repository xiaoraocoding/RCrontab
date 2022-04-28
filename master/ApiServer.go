package master

import (
	"RCrontab/common"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)


func InitRouter(r *gin.Engine)  {
	r.GET("/index",HandleIndex)
	r.POST("/job/save",HanleJobSave)
	r.POST("/job/delete",HandleDeleteJob)
	r.GET("/job/list",HandleList)
	r.POST("/job/kill",HandleKill)
	r.GET("/job/log",HandleLog)
	r.GET("/worker/list",HandleWordList)

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
		"errno": 0,
		"msg" : "success",
		"data":oldJob,
	})
}

func HandleDeleteJob(c *gin.Context) {
	name := c.PostForm("name")

	oldJob,err:= G_JobMgr.Delete(name)
	if err != nil {
		fmt.Println("delete failed,err:",err)
	}
	c.JSON(200,gin.H{
		"errno": 0,
		"msg" : "success",
		"data":oldJob,
	})


}

func HandleList(c *gin.Context) {
	list,err := G_JobMgr.ListJobs()
	if err != nil {
		fmt.Println("list failed err:",err)
	}
	c.JSON(200,gin.H{
		"errno": 0,
		"msg" : "success",
		"data":list,
	})


}


func HandleKill(c *gin.Context) {
	name := c.PostForm("name")
	err := G_JobMgr.Kill(name)
	if err != nil {
		fmt.Println("kill filed,err:",err)
	}
	c.JSON(200,gin.H{
		"errno": 0,
		"msg" : "success",
		"data":nil,
	})



}

func HandleIndex(c *gin.Context) {
	c.HTML(200,"index.html","")
}


func HandleLog(c *gin.Context) {
	// 这里传来的数据是 /job/log?name=jon01&skip=0&limit=10
	name := c.Query("name")
	skip := c.Query("skip")
	limit := c.Query("limit")

	sk,_ := strconv.Atoi(skip)
	lim,_ := strconv.Atoi(limit)
	var logArr []common.JobLog
	var err error

	if logArr,err = G_logMgr.ListLog(name,sk,lim); err != nil {
		fmt.Println("listLog failed err:",err)
		return
	}
	fmt.Println(logArr)


	c.JSON(200,gin.H{
		"errno": 0,
		"msg" : "success",
		"data":&logArr,
	})

}

func HandleWordList(c *gin.Context) {
	s,err := G_workerMgr.ListWorkers()
	if err != nil {
		fmt.Println("listWork",err)
	}
	c.JSON(200, gin.H{
		"errno": 0,
		"msg" : "success",
		"data":s,
	})

}
