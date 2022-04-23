package main

import (
	"RCrontab/worker"
	"time"
)

func main() {
	worker.InitSchedule()
	worker.InitJobMgr()

	worker.InitLogSink()
    worker.InitRegister()



	for {
		time.Sleep(1*time.Second)
	}



}