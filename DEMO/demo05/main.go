package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var config clientv3.Config
	var client *clientv3.Client
	var err error
	var lease clientv3.Lease
	var leastRes *clientv3.LeaseGrantResponse

	//etcd的客户端链接配置
	config = clientv3.Config{
		Endpoints: []string{},
		DialTimeout: 5*time.Second,

	}

	if client,err = clientv3.New(config) ; err != nil {
		fmt.Println(err)
		return
	}
    //创建一个一个租约
	lease = clientv3.NewLease(client)

	//申请一个10秒的租约
	leastRes,err = lease.Grant(context.TODO(),10)
	if err != nil {
		fmt.Println(err)
	}
	leastId := leastRes.ID

	//一直续租
	keepChan,_ := lease.KeepAlive(context.TODO(),leastId)

	go func() {
		select {
		case res := <- keepChan:
			if res == nil {
				fmt.Println("因为出现的某些问题，此时的租约已经过期")
			}else {
				fmt.Println("收到续租，续租Id：",res.ID)
			}
			return
		}
	}()


	kv := clientv3.NewKV(client)
	r,_ := kv.Put(context.TODO(),"name","10s_failed",clientv3.WithLease(leastId))
	fmt.Println(r.Header.Revision)


	time.Sleep(20*time.Second)







}
