package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var err error
	var expr *cronexpr.Expression
	var now time.Time
	if expr,err = cronexpr.Parse("*/5 * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	now = time.Now()
	nexTime := expr.Next(now)
	fmt.Println(now,nexTime)




}
