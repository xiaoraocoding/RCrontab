package worker

import "RCrontab/common"

type Scheduler struct {
	jobEventChan chan *common.JobEvent //etcd的任务队列
	jobPlanTable  map[string]*common.JobSchedulePlan
}

var W_Scheduler *Scheduler

func (schedule *Scheduler) handleJobEvent(jobEvent *common.JobEvent) {
	switch jobEvent.EventType {
	case 1: //修改

	case 2: //删除

	}


}
func (schedule *Scheduler) scheduleLoop() {



}

func InitSchedule() {
	W_Scheduler = &Scheduler{
		jobEventChan: make(chan *common.JobEvent,1000),
	}
	go W_Scheduler.scheduleLoop()
}

func (schedule *Scheduler)PushJobEvent(jobEvent *common.JobEvent) {
	schedule.jobEventChan <- jobEvent
}



