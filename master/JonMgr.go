package master

import (
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
