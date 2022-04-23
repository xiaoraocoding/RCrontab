package common

type JobLog struct {
	JobName string `bson:"jobName" json:"jobName"`
	Command string `bson:"command" json:"command"`
	Err string `bson:"err" json:"err"`
	StartTime int64 `bson:"startTime" json:"startTime"`
	EndTime int64 `bson:"endTime" json:"endTime"`
	OutPut string `bson:"output" json:"output"`
	PlanTime int64 `bson:"planTime" json:"planTime"`
	ScheduleTime int64 `json:"scheduleTime" bson:"scheduleTime"`
}

//当我们从chan中读取到log后，我们可以先放到logBatch里面，等暂留一定的数据量后再发送
type LogBatch struct {
	Logs []interface{}
}


type JobLogFilter struct {
	JobName string `bson:"jobName"`
}

// 任务日志排序规则
type SortLogByStartTime struct {
	SortOrder int `bson:"startTime"`	// {startTime: -1}
}


