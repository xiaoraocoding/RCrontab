package main

import (
	"RCrontab/worker"
	"time"
)

func main() {
	worker.InitJobMgr()


	worker.InitSchedule()

	for {
		time.Sleep(1*time.Second)
	}



}