package main

import (
	"job-hunter/engine"
	parser "job-hunter/parser/shixiseng"
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
	e.Run(engine.Request{
		Url:       "https://www.shixiseng.com/interns?type=intern&city=%E5%85%A8%E5%9B%BD",
		ParseFunc: parser.ParseJobList,
	})
	//res, err := http.Get("https://www.shixiseng.com/intern/inn_n7xuu1c0hizv?pcm=pc_SearchList")
	//if err != nil {
	//	panic(err)
	//}
	//data := parser.ParseJob(res.Body, "")
	//fmt.Printf("%v",data.Items[0])
	//low, high := libs.ParseSalary("\uEB5B\uF0E5-\uE0D8\uF0E5\uF0E5/天", shixiseng.ConvertNumber)
	//fmt.Println(low, high)

}
