package main

import (
	"RCrontab/worker"
	"time"
)

func main() {
	worker.InitSchedule()
	worker.InitJobMgr()




	for {
		time.Sleep(1*time.Second)
	}



}