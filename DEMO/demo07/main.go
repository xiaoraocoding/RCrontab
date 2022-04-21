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



	//etcd的客户端链接配置
	config = clientv3.Config{
		Endpoints:   []string{},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)
	putOp := clientv3.OpPut("test/02","1")

	res,err := kv.Do(context.TODO(),putOp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Put().Header.Revision)

	getOp := clientv3.OpGet("test/02")

	resGet,err := kv.Do(context.TODO(),getOp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("版本:",resGet.Get().Kvs[0].ModRevision)
	fmt.Println("value",resGet.Get().Kvs[0].Value)



}