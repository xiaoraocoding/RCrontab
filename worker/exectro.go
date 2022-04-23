package worker

import (
	"RCrontab/common"
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

type Executor struct {

}

var W_Executor *Executor


func (executor *Executor) ExecuteJob(job *common.JobExecuteInfo) {
	go func() {
		//上锁之前进行随机的睡眠
		time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
		joblock := W_JobMgr.CreateLock(job.Job.Name)  //创建一个锁
      var result *common.JobExecuteResult
		 err := joblock.TryLock()
		 defer joblock.Unlock()
		 if err != nil {
			 start := time.Now()
			fmt.Println("这里的上锁失败了",err)

			result = &common.JobExecuteResult{
				ExecuteInfo: job,
				Err: err,
				StartTime: start,
			}
		}else {
			start := time.Now()
			cmd := exec.CommandContext(job.CanCtx,"/bin/bash","-c",job.Job.Command)
			output,err := cmd.CombinedOutput()
			if err  != nil {
				fmt.Println("server run command failed",err)
				return
			}
			end := time.Now()

			result = &common.JobExecuteResult{
				ExecuteInfo: job,
				Output: output,
				Err: err,
				StartTime: start,
				EndTime: end,
			}


		}
		W_Scheduler.PushJobResult(result)
		}()



}
