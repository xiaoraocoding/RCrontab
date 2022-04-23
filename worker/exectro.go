package worker

import (
	"RCrontab/common"
	"context"
	"fmt"
	"os/exec"
	"time"
)

type Executor struct {

}

var W_Executor *Executor


func (executor *Executor) ExecuteJob(job *common.JobExecuteInfo) {
	go func() {
		start := time.Now()
		cmd := exec.CommandContext(context.TODO(),"/bin/bash","-c",job.Job.Command)
		output,err := cmd.CombinedOutput()
		if err  != nil {
			fmt.Println("server run command failed",err)
			return
		}
		end := time.Now()

		result := &common.JobExecuteResult{
			ExecuteInfo: *job,
			Output: output,
			Err: err,
			StartTime: start,
			EndTime: end,
		}
		W_Scheduler.PushJobResult(result)
	}()

}
