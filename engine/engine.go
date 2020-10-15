package engine

import (
	"strconv"
)

type Engine struct {
	WorkerLimit int
	// channel for output item
	ItemChan  chan Item
	Scheduler Scheduler
	PageLimit int
}
type Scheduler interface {
	MakeWorkerChan() chan Request
	SubmitRequest(Request)
	SubmitWorker(chan Request)
	Run()
}

func (e Engine) Run(seed Request, isPageInc bool) {
	out := make(chan ParseResult)
	var requests []Request

	e.Scheduler.Run()

	for i := 0; i < e.WorkerLimit; i++ {
		createWorker(e.Scheduler, e.Scheduler.MakeWorkerChan(), out)
	}
	// generate request by page
	if isPageInc {
		requestGenerator := GenerateUrlPage(seed)
		for i := 0; i < e.PageLimit; i++ {
			requests = append(requests, requestGenerator())
		}
	} else {
		requests = append(requests, seed)
	}

	for _, r := range requests {
		e.Scheduler.SubmitRequest(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				// push to item channel to saver
				e.ItemChan <- item
			}()
		}
		for _, r := range result.Requests {
			e.Scheduler.SubmitRequest(r)
		}
	}
}

func createWorker(s Scheduler, in chan Request, out chan ParseResult) {
	go func() {
		for {
			// submit available worker
			s.SubmitWorker(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func GenerateUrlPage(r Request) func() Request {
	page := 0
	return func() Request {
		page++
		return Request{r.Url + "&page=" + strconv.Itoa(page), r.ParseFunc}
	}
}
