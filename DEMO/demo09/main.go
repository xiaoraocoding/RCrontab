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
	fmt.Println(colletion)

	//record1 := &logRecord{
	//	JobName: "job02",
	//	Command: "echo hello",
	//	Err: "",
	//	Context: "hello",
	//	TimePoint:TimePoint{
	//		StartTime: time.Now().Unix(),
	//		EndTime: time.Now().Unix()+10,
	//	},
	//}
	//record2 := &logRecord{
	//	JobName: "job023",
	//	Command: "echo hello",
	//	Err: "",
	//	Context: "hello",
	//	TimePoint:TimePoint{
	//		StartTime: time.Now().Unix(),
	//		EndTime: time.Now().Unix()+10,
	//	},
	//}
	record3 := &logRecord{
		JobName: "job04",
		Command: "echo golang",
		Err: "",
		Context: "golang",
		TimePoint:TimePoint{
			StartTime: time.Now().Unix(),
			EndTime: time.Now().Unix()+10,
		},
	}

	log := []interface{}{record3}


	res,err := colletion.InsertMany(context.TODO(),log)
    if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.InsertedIDs)






	//3 选择表
}


