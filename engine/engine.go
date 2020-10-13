package engine

type Engine struct {
	WorkerLimit int
	ItemChan    chan Item
	Scheduler   Scheduler
}
type Scheduler interface {
	SubmitRequest(Request)
	SubmitWorker(chan Request)
	Run()
}

func (e Engine) Run(requests ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerLimit; i++ {
		createWorker(e.Scheduler, out)
	}

	for _, r := range requests {
		e.Scheduler.SubmitRequest(r)
	}

}

func createWorker(s Scheduler, out chan ParseResult) {
	in := make(chan Request)
	go func() {
		for {
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
