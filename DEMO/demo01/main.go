package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	//var outPut []byte
	//var err error
	//
	////cmd命令
	//cmd := exec.Command("/bin/bash","-c","sleep 8;go run /Users/raoraoningkang/Desktop/MyProject/RCrontab/text/main.go")
	//
	////子进程运行cmd，读取出子进程的数据
	//if outPut,err = cmd.CombinedOutput() ; err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(outPut))

	var outPut []byte
	var err error
	var ctx context.Context
	var cancelFunc context.CancelFunc
	var res chan *result   //这里感觉放指针的话会更小点

	ctx,cancelFunc = context.WithCancel(context.TODO())
	res = make(chan *result,2)

    go func() {
		//cmd命令
		cmd := exec.CommandContext(ctx,"/bin/bash","-c","sleep 3;go run /Users/raoraoningkang/Desktop/MyProject/RCrontab/text/main.go")

		//子进程运行cmd，读取出子进程的数据
		if outPut,err = cmd.CombinedOutput() ; err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(outPut))

		res <- &result{
			err: err,
			res: string(outPut),
		}
	}()

	time.Sleep(1*time.Second)
	//等待一秒后取消协程
	cancelFunc()


	fmt.Println("hello golang")

}

type result struct {
	err error
	res string
}
