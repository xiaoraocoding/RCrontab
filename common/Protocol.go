package common

//任务
type Job struct {
	Name string `json:"name"`
	Command string `json:"command"`   //任务的命令
	CronExpr string `json:"cron_expr"`  //定时
}


