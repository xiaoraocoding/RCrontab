package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type TimeBeforeCond struct {

	Before int64 `bson:"$lt"`
}

type DeleteCond struct {
	beforeCond TimeBeforeCond `bson:"timePoint.startTime"`

}
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("m7"))
	if err != nil {
		fmt.Println(err)
	}


	data := client.Database("cron")

	colletion := data.Collection("log")

	delCond := &DeleteCond{
		beforeCond: TimeBeforeCond{Before: time.Now().Unix()},
	}

	colletion.DeleteMany(context.TODO(),delCond)





}