package main

import (
	"context"
	"job-hunter/engine"
	"job-hunter/parser/shixiseng"
	"job-hunter/persist"
	"job-hunter/scheduler"
)

func main() {
	itemChan, client, err := persist.Saver("shixiseng")
	if err != nil {
		panic(err)
	}
	e := engine.Engine{
		Scheduler:   &scheduler.Scheduler{},
		WorkerLimit: 20,
		ItemChan:    itemChan,
		PageLimit:   500,
	}
	e.Run(engine.Request{
		Url:       "https://www.shixiseng.com/interns?type=intern&city=%E5%85%A8%E5%9B%BD",
		ParseFunc: shixiseng.ParseJobList,
	}, true)

	client.Disconnect(context.TODO())
}