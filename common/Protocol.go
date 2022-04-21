package common

import (
	"encoding/json"
	"strings"
)



//任务
type Job struct {
	Name string `json:"name"`
	Command string `json:"command"`   //任务的命令
	CronExpr string `json:"cron_expr"`  //定时
}

type JobEvent struct {
	EventType int  //此处分两种，一种是delete，一种是put，也就是删除和修改
	job *Job
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


