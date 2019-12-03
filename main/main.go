package main

import (
	"code.byted.org/gopkg/pkg/log"
	"fmt"
	. "local/go_tool_dev/work_pool"
	"time"
)

func main() {
	log.Infoln("....... start")
	//创建一个Task
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	//创建一个协程池,最大开启5个协程worker
	p := NewWorkPool(5)
	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for {
			p.Job <- t
		}
	}()

	//启动协程池p
	p.Run()
	time.Sleep(10 * time.Second)

}