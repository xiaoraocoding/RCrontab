package worker

import (
	"RCrontab/common"
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

type JobMgr struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
	watch clientv3.Watcher
}

var W_JobMgr *JobMgr

func InitJobMgr() {
	conf := clientv3.Config{
		Endpoints: []string{"47.96.162.129:2379"},
		DialTimeout: 5*time.Second,
	}
	client,err := clientv3.New(conf)
	if err != nil {
		fmt.Println("clientv3.New() failed err:",err)
	}
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)
	watch := clientv3.NewWatcher(client)

	W_JobMgr = &JobMgr{
		client: client,
		kv: kv,
		lease: lease,
		watch: watch,
	}
	fmt.Println(W_JobMgr)
}

//监听任务的变化
func (jobMgr *JobMgr) watchJobs() (err error) {
	var job *common.Job

	//1 知道了此时目录下的所有的任务，得到当前节点的revision
	getRes,err := jobMgr.kv.Get(context.TODO(),"/cron/jobs/",clientv3.WithPrefix())
	if err != nil {
		fmt.Println("watch work filed,err:",err)
	}
	for _,v := range getRes.Kvs {
		if job,err = common.UnpackJob(v.Value);err ==nil {
			jobEvent := common.BuildJobEvent(1,job)

		}

	}
	//2 从这个端口监听 revision的变化
	go func() {
		watchStartRevison := getRes.Header.Revision + 1
		watchChan := jobMgr.watch.Watch(context.TODO(),"/cron/jobs/",clientv3.WithRev(watchStartRevison))
		for watchRes := range watchChan {
			for _,watchEvent := range watchRes.Events {
				switch watchEvent.Type{
				case mvccpb.PUT:
					//出现了异常
					if job,err = common.UnpackJob(watchEvent.Kv.Key);err != nil {
						continue
					}
					//构建一个更新的Event
					jobEvent := common.BuildJobEvent(1,job)

					//任务保存
				case mvccpb.DELETE:
					//任务删除
					jobName:= common.ExtrateJonName(string(watchEvent.Kv.Value))
					job = &common.Job{Name: jobName}
					//构建一个删除的Event
					jobEvent := common.BuildJobEvent(2,job)



				
				}

			}
		}

	}()
}