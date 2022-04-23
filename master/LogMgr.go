package master

import (
	"RCrontab/common"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// mongodb日志管理
type LogMgr struct {
	client *mongo.Client
	logCollection *mongo.Collection
}

var (
	G_logMgr *LogMgr
)

func InitLogMgr() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://47.96.162.129:27017"))
	if err != nil {
		fmt.Println(err)
	}
	G_logMgr = &LogMgr{
		client:        client,
		logCollection: client.Database("cron").Collection("log"),
	}
	return err
}


// 查看任务日志
func (logMgr *LogMgr) ListLog(name string, skip int, limit int) (logArr []common.JobLog, err error){
	var (
		filter *common.JobLogFilter
		logSort *common.SortLogByStartTime
		cursor *mongo.Cursor
		jobLog *common.JobLog
	)

	// len(logArr)
	logArr = make([]common.JobLog, 0)

	// 过滤条件
	filter = &common.JobLogFilter{JobName: name}

	// 按照任务开始时间倒排
	logSort = &common.SortLogByStartTime{SortOrder: -1}

	// 查询
	opts := &options.FindOptions{}
	opts.SetSkip(int64(skip))
	opts.SetLimit(int64(limit))
    opts.SetSort(logSort)
	if cursor, err = logMgr.logCollection.Find(context.TODO(), filter, opts); err != nil {
		return
	}
	// 延迟释放游标

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		jobLog = &common.JobLog{}

		// 反序列化BSON
		if err = cursor.Decode(jobLog); err != nil {
			continue // 有日志不合法
		}


		logArr = append(logArr, *jobLog)

	}
	return
}

