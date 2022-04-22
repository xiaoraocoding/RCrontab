package worker

import (
	"RCrontab/common"
	"fmt"
	"time"
)

type Scheduler struct {
	jobEventChan chan *common.JobEvent //etcd的任务队列
	jobPlanTable  map[string]*common.JobSchedulePlan
	jobExecutingTable map[string]*common.JobExecuteInfo
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
	after := schedule.TrySchedule()
	timerA := time.NewTimer(after)

	for {
		select {
		case jobEvent := <- schedule.jobEventChan:
			schedule.handleJobEvent(jobEvent)
			case <- timerA.C :

		}
		after = schedule.TrySchedule()
		timerA.Reset(after)

		}
	}


//其实这些只是为了做到同步，也就是让数据做到和ecd里面的数据同步
func InitSchedule() {
	W_Scheduler = &Scheduler{
		jobEventChan: make(chan *common.JobEvent,1000),
		jobPlanTable: make(map[string]*common.JobSchedulePlan),
		jobExecutingTable: make(map[string]*common.JobExecuteInfo),
	}
	go W_Scheduler.scheduleLoop()

}

func (schedule *Scheduler)PushJobEvent(jobEvent *common.JobEvent) {
	schedule.jobEventChan <- jobEvent
}

//个人理解，就是在我们实际的环境中，并不是每秒都在执行任务，很多时候都是在等待任务的时间到来
//那么这个时候，我们就可以让cpu进行睡眠，实现的话，经过我们之前的包装，可以很简单的实现出来
func (shedule *Scheduler) TrySchedule()time.Duration {
	now := time.Now()
	var near *time.Time
	if len(shedule.jobPlanTable) == 0 {  //也就是当没有任务的时候，这个时候其实无所谓了
		return  1 * time.Second
	}

	for _,jobPlan := range shedule.jobPlanTable {
		//这里需要注意的就是 这里有可能执行有可能不执行，原因是因为可能你的任务的执行时间长，那么
		//到了这次的执行时间，你还是没有执行完，那就不执行来
		if jobPlan.NextTime.Before(now) || jobPlan.NextTime.Equal(now) {
			jobPlan.NextTime = jobPlan.Expr.Next(now)  //下一次更新的时间
			shedule.TryStartJob(jobPlan)
		}
		if near == nil || jobPlan.NextTime.Before(*near) {
			near = &jobPlan.NextTime
		}
	}
	scheduleAfter := near.Sub(now)
	return scheduleAfter
}

//尝试执行任务(这个主要解决的就是当我们调用的服务时间要很多，但是每一次间隔的时间短，这就不需要执行了)
func (schedule *Scheduler) TryStartJob (jobPlan *common.JobSchedulePlan) {
     var jobExecultInfo *common.JobExecuteInfo
	 var jobExing bool
	 //如果还在执行，直接返回，说明不用执行了
	 if jobExecultInfo,jobExing = schedule.jobExecutingTable[jobPlan.Job.Name];jobExing{
		 fmt.Println("已经在执行了----")
         return
	 }

	 jobExecultInfo = common.BuildJobExecuteInfo(jobPlan)

	 schedule.jobExecutingTable[jobPlan.Job.Name] = jobExecultInfo

	 //todo:  此处应该写并发的执行任务
	 fmt.Println("执行任务",jobPlan.Job.Name,jobExecultInfo.PlanTime,jobExecultInfo.RealTime)



}



