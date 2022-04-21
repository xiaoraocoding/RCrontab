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
	var kv clientv3.KV
	var res *clientv3.PutResponse
	var getRes *clientv3.GetResponse

	//etcd的客户端链接配置
	config = clientv3.Config{
		Endpoints: []string{"47.96.162.129:2379"},
		DialTimeout: 5*time.Second,

	}

	if client,err = clientv3.New(config) ; err != nil {
		fmt.Println(err)
		return
	}

	kv =  clientv3.NewKV(client)
	res,err = kv.Put(context.TODO(),"name","xiaoraocasdasdng",clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	getRes,err = kv.Get(context.TODO(),"/cron/name/",clientv3.WithPrefix())
    if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getRes.Kvs)
}
