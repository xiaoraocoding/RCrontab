package common

import (
	"encoding/json"
	"fmt"
	"github.com/gorhill/cronexpr"
	"strings"
	"time"
)



//任务
type Job struct {
	Name string `json:"name"`
	Command string `json:"command"`   //任务的命令
	CronExpr string `json:"cron_expr"`  //定时
}

type JobEvent struct {
	EventType int  //此处分两种，一种是delete，一种是put，也就是删除和修改
	Job *Job
}

//任务执行计划
type JobSchedulePlan struct {
	Job *Job
	Expr *cronexpr.Expression
	NextTime time.Time  //下一次的调度时间
}

func UnpackJob(value []byte)(res *Job,err error){
	job := &Job{}

	err = json.Unmarshal(value,job)
	if err != nil {
		return nil,err
	}
	res = job
	return

}

func ExtrateJonName(jobKey string)(string) {

	return strings.TrimPrefix(jobKey,"/cron/jobs/")
}

//构建事件的变化
func BuildJobEvent(eventType int,job *Job) *JobEvent{
	return &JobEvent{
		eventType,
		job,
	}
}

//构建任务计划
func BuildJobSchedulePlan(job *Job)*JobSchedulePlan {
	//解析cron表达式
	expr,err := cronexpr.Parse(job.CronExpr)
	if err != nil {
		fmt.Println("parse failed",err)
	}

	jobSchedulePlan:= &JobSchedulePlan{
		job,
		expr,
		expr.Next(time.Now()),

	}
	return jobSchedulePlan
}


