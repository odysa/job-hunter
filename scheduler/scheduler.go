package scheduler

import (
	"job-hunter/engine"
	"job-hunter/libs"
)

type Scheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
	RequestQ    libs.QueueInterface
	WorkerQ     libs.QueueInterface
}

func (s *Scheduler) SubmitRequest(request engine.Request) {
	s.RequestChan <- request
}

func (s *Scheduler) SubmitWorker(worker chan engine.Request) {
	s.WorkerChan <- worker
}

func (s *Scheduler) Run() {
	s.RequestChan = make(chan engine.Request)
	s.WorkerChan = make(chan chan engine.Request)
	go func() {
	}()
}
