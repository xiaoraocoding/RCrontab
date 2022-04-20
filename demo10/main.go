package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


type TimePoint struct {
	StartTime int64 `bson:"start_time"`
	EndTime int64  `bson:"end_time"`
}

type logRecord struct {
	JobName string  `bson:"job_name"`//任务名称
	Command string `bson:"command"`//脚本命令
	Err string      `bson:"err"`//错误提示
	Context string `bson:"context"`//脚本输出
	TimePoint `bson:"time_point"`

}

type FindByJonName struct {
	JobName string  `bson:"job_name"`//任务名称

}
//链接服务器上的mogodb 以及实现CRUD
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://47.96.162.129:27017"))
	if err != nil {
		fmt.Println(err)
	}


	data := client.Database("cron")

	colletion := data.Collection("log")

	cond := &FindByJonName{JobName: "job04"}


	//设置边界
	opts := options.Find().SetSkip(0)
	opts.SetLimit(2)



	cur,err := colletion.Find(context.TODO(),cond,opts)
	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		record := &logRecord{}
		err := cur.Decode(record)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*record)

	}




}

