package master

import (
	"RCrontab/common"
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type JobMgr struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}

var G_JobMgr *JobMgr

func InitJobMgr() {
	conf := clientv3.Config{
		Endpoints: []string{Conf.Etcd},
		DialTimeout: 5*time.Second,
	}
	client,err := clientv3.New(conf)
	if err != nil {
		fmt.Println("clientv3.New() failed err:",err)
	}
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)

	G_JobMgr = &JobMgr{
		client: client,
		kv: kv,
		lease: lease,
	}
	fmt.Println(G_JobMgr)
}


//保存任务(如果之前有老的值的话，那就返回出来这个老的值)
func (jobMgr *JobMgr) SaveJob(job *common.Job)(oldJob *common.Job,err error) {
	jobKey := "/cron/jobs/" + job.Name

	jobValue,err := json.Marshal(job)
	if err != nil {
		fmt.Println("json.Marshal failed,err:",err)
		return nil,err
	}
	putRes,err := jobMgr.kv.Put(context.TODO(),jobKey,string(jobValue),clientv3.WithPrevKV())
	if err != nil {
		fmt.Println("save etcd failed, err:",err)
		return nil,err
	}

	if putRes.PrevKv != nil {
		json.Unmarshal(putRes.PrevKv.Value,&oldJob)
		err = nil
		return
	}
    return

}

func (jobMgr *JobMgr) Delete(name string) (oldJob *common.Job,err error) {
	jobKey := "/cron/jobs/" + name

	deleteRes,err := jobMgr.client.Delete(context.TODO(),jobKey,clientv3.WithPrevKV())
	if err != nil {
		fmt.Println("Delete failed err:",err)
		return nil,err
	}

	if deleteRes.PrevKvs != nil {
		json.Unmarshal(deleteRes.PrevKvs[0].Value,&oldJob)
		err = nil
		return
	}
	return

}


//列举所以的任务
func (jobMgr *JobMgr) ListJobs() (jobList []*common.Job,err error) {
	jobKey := "/cron/jobs/"

	getRes,err := jobMgr.client.Get(context.TODO(),jobKey,clientv3.WithPrefix())
	if err != nil {
		return nil,err
	}
	jobList = make([]*common.Job,0)
	for _,v := range getRes.Kvs {
		job := common.Job{}
		json.Unmarshal(v.Value,&job)
		jobList = append(jobList,&job)
	}
	err = nil
	return

}
