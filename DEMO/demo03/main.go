package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr *cronexpr.Expression
	nextTime time.Time
}
func main() {
	var c1 *CronJob
	var c2 *CronJob

	shedule := make(map[string]*CronJob)

	var err error
	var expr1 *cronexpr.Expression
	var now1 time.Time
	if expr1,err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	now1 = time.Now()
	c1 = &CronJob{
		expr: expr1,
		nextTime: expr1.Next(now1),
	}


	var expr2 *cronexpr.Expression
	var now2 time.Time
	if expr2,err = cronexpr.Parse("*/3 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	now2 = time.Now()
	c2 = &CronJob{
		expr: expr2,
		nextTime: expr2.Next(now2),
	}
	shedule["job1"] = c1
	shedule["job2"] = c2

	for {
		now := time.Now()
		for jobName,cronJob := range shedule {
			if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
				//这里的用法的原因是使用goroutine可以不用等待，如果不使用，那会进入等待
				go func(job string) {
					fmt.Println(job)
				}(jobName)
				cronJob.nextTime = cronJob.expr.Next(now)  //将时间换为当前的值

			}
		}

		select {   //这里停止100ms的原因是防止将服务器的资源占满，最好还是停顿几ms
		case <- time.NewTimer(100*time.Millisecond).C:
		}

	}
	time.Sleep(100*time.Second)   //这里睡眠100秒
}
