package main

import (
	"RCrontab/worker"
	"time"
)

func main() {
	worker.InitJobMgr()

	worker.W_JobMgr.WatchJobs()

	for {
		time.Sleep(1*time.Second)
	}



}