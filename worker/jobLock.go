package worker

import "github.com/coreos/etcd/clientv3"

//分布式锁
type JobLock struct {
	JobName string //锁住服务的名字
	KV clientv3.KV
	Lease clientv3.Lease //服务的租约
}

func InitJobLock(name string,kv clientv3.KV,lease clientv3.Lease) *JobLock {
	return &JobLock{
		name,
		kv,
		lease,
	}

}