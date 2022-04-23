package worker

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

//分布式锁
type JobLock struct {
	JobName string //锁住服务的名字
	KV clientv3.KV
	Lease clientv3.Lease //服务的租约
	CancelFunc context.CancelFunc  //用来终止自动续租，这里的主要的作用就是解锁
	LeaseId clientv3.LeaseID
	IsLocked bool //是否上锁成功
}

func InitJobLock(name string,kv clientv3.KV,lease clientv3.Lease) *JobLock {
	return &JobLock{
		name,
		kv,
		lease,
		nil,
		0,
		false,
	}

}


//这里是上锁的具体的实现，其实上锁的步骤可以细分
func (joblock *JobLock)TryLock() error {

	//1 创建租约
	leasePes,err := joblock.Lease.Grant(context.TODO(),5)
	if err != nil {
		fmt.Println("grant failed err:",err)
		return err
	}
	//2 自动续租
	canctx,cancelFunc := context.WithCancel(context.TODO())

	leaseId := leasePes.ID
	keepRes,err := joblock.Lease.KeepAlive(canctx,leaseId)
    if err != nil {
		fmt.Println("keepAlive failed err:",err)
		cancelFunc()
		joblock.Lease.Revoke(canctx,leaseId)
		return err
	}
	//3 处理

	go func() {
		for {
			select {
			case keepChan := <- keepRes:
				if keepChan == nil {
					goto END
				}
			}
		}
		END:
	}()

	//4 创建事务txn
	txn := joblock.KV.Txn(context.TODO())

	lockKey := "/cron/lock/" + joblock.JobName

	//事务抢锁
	txn.If(clientv3.Compare(clientv3.CreateRevision(lockKey),"=",0)).
   Then(clientv3.OpPut(lockKey,"",clientv3.WithLease(leaseId))).Else(clientv3.OpGet(lockKey))

	//提交事务
	txnRes,err := txn.Commit()
	if err != nil {
		fmt.Println("txn commit failed err:",err)
		cancelFunc()
		joblock.Lease.Revoke(canctx,leaseId)  //这里的锁都需要进行释放(这里的key，直接被删除了)
        return err
	}

	//查看抢锁是否成功
	if !txnRes.Succeeded {
		fmt.Println("锁被占用")
		cancelFunc()
		joblock.Lease.Revoke(canctx,leaseId)
		return err
	}

	//此时抢锁已经成功

	joblock.CancelFunc = cancelFunc
	joblock.LeaseId = leaseId
	joblock.IsLocked = true

	return nil

}

//释放锁
func (jobLock *JobLock) Unlock() {
	if jobLock.IsLocked {
		jobLock.CancelFunc()
		jobLock.Lease.Revoke(context.TODO(),jobLock.LeaseId)
	}
}
