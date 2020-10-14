package scheduler

import (
	"job-hunter/engine"
	"job-hunter/libs"
)

type Scheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (s *Scheduler) SubmitRequest(request engine.Request) {
	s.RequestChan <- request
}

func (s *Scheduler) SubmitWorker(worker chan engine.Request) {
	s.WorkerChan <- worker
}

func (s *Scheduler) MakeWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *Scheduler) Run() {
	s.RequestChan = make(chan engine.Request)
	s.WorkerChan = make(chan chan engine.Request)
	go func() {
		// queue
		var requestQ libs.Queue
		var workerQ libs.Queue
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// if not empty got active item
			if !requestQ.IsEmpty() && !workerQ.IsEmpty() {
				activeRequest = requestQ.Top().(engine.Request)
				activeWorker = workerQ.Top().(chan engine.Request)
			}
			select {
			case r := <-s.RequestChan:
				requestQ.Push(r)
			case w := <-s.WorkerChan:
				workerQ.Push(w)
			case activeWorker <- activeRequest:
				requestQ.Pop()
				workerQ.Pop()
			}
		}
	}()
}
