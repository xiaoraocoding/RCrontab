package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

func main() {
	var config clientv3.Config
	var client *clientv3.Client
	var err error
	var kv clientv3.KV
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

	go func() {
		for {
			kv.Put(context.TODO(),"test/01","a")
			kv.Delete(context.TODO(),"test/01")
			time.Sleep(1*time.Second)
		}
	}()

	getRes,err = kv.Get(context.TODO(),"test/01")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(getRes.Kvs) != 0 {
		fmt.Println(string(getRes.Kvs[0].Value))
	}

	watchId := getRes.Header.Revision + 1

	watcher := clientv3.NewWatcher(client)

	watchChan := watcher.Watch(context.TODO(),"test/01",clientv3.WithRev(watchId))

	for wathRes := range watchChan {
		for _,e := range wathRes.Events {
			switch e.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:" + string(e.Kv.Value),"version",e.Kv.ModRevision,e.Kv.CreateRevision)
			case mvccpb.DELETE:
				fmt.Println("删除",e.Kv.ModRevision)

			}
		}
	}


}
