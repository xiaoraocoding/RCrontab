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
	jobPlantable := common.BuildJobSchedulePlan(jobEvent.Job)
	schedule.jobPlanTable[jobEvent.Job.Name] = jobPlantable

	case 2: //删除
	_,ok := schedule.jobPlanTable[jobEvent.Job.Name]
	if ok {
		delete(schedule.jobPlanTable,jobEvent.Job.Name)
	}
	}


}
func (schedule *Scheduler) scheduleLoop() {



}

//其实这些只是为了做到同步，也就是让数据做到和ecd里面的数据同步
func InitSchedule() {
	W_Scheduler = &Scheduler{
		jobEventChan: make(chan *common.JobEvent,1000),
		jobPlanTable: make(map[string]*common.JobSchedulePlan),
	}
	go W_Scheduler.scheduleLoop()

}

func (schedule *Scheduler)PushJobEvent(jobEvent *common.JobEvent) {
	schedule.jobEventChan <- jobEvent
}



