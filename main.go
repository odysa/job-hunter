package main

import (
	"job-hunter/engine"
	"job-hunter/parser/zhilian"
	"job-hunter/persist"
	"job-hunter/scheduler"
)

func main() {
	itemChan, err := persist.Saver("test")
	if err != nil {
		panic(err)
	}
	e := engine.Engine{
		Scheduler:   &scheduler.Scheduler{},
		WorkerLimit: 10,
		ItemChan:    itemChan,
		PageLimit:   10,
	}
	//e.Run(engine.Request{
	//	Url:       "https://www.shixiseng.com/interns?type=intern&city=%E5%85%A8%E5%9B%BD",
	//	ParseFunc: parser.ParseJobList,
	//})
	e.Run(engine.Request{
		Url:       "https://www.zhaopin.com/citymap",
		ParseFunc: zhilian.ParseCity,
	}, false)
}
