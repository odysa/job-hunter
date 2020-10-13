package scheduler

import (
	"job-hunter/engine"
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

func (s Scheduler) Run() {

}
