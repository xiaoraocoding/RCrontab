package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

//这里的demo主要是为了实现分布式的锁

func main() {
	var config clientv3.Config
	var client *clientv3.Client
	var err error
	var kv clientv3.KV
	var lease clientv3.Lease
	var leastRes *clientv3.LeaseGrantResponse

	//etcd的客户端链接配置
	config = clientv3.Config{
		Endpoints:   []string{"47.96.162.129:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}


	lease = clientv3.NewLease(client)
	//下面的1 2 3 就是上锁的步骤
	//1 上锁，创建自动租约，自动续租，拿租约去占key

	//申请一个10秒的租约
	leastRes, err = lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println(err)
	}
	leastId := leastRes.ID

	ctx,cancel := context.WithCancel(context.TODO())

    defer cancel()  //这里是为了防止意外，不退出
	defer lease.Revoke(context.TODO(),leastId)  //没有租约，直接放开key

	keepChan, _ := lease.KeepAlive(ctx, leastId)

	go func() {
		select {
		case res := <-keepChan:
			if res == nil {
				fmt.Println("因为出现的某些问题，此时的租约已经过期")
			} else {
				fmt.Println("收到续租，续租Id：", res.ID)
			}
			return
		}
	}()

	kv = clientv3.NewKV(client)

	//定义事务
	tnx := kv.Txn(context.TODO())
	tnx.If(clientv3.Compare(clientv3.CreateRevision("/demo11"),"=",0)).
		Then(clientv3.OpPut("/demo11","22",clientv3.WithLease(leastId))).
		Else(clientv3.OpGet("/demo11"))

	 //提交事务
	res,err := tnx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

	if !res.Succeeded {
		fmt.Println("抢锁失败",string(res.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}



	//2 处理业务  （运行到此处，说明已经得到了锁，此时是安全的）

	fmt.Println("处理任务")
	time.Sleep(10*time.Second)



}

