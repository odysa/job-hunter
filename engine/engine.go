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

}

func createWorker(s Scheduler, out chan ParseResult) {
	in := make(chan Request)
	for {
		s.SubmitWorker(in)
		request := <-in
		result := worker(request)
		out <- result
	}
}
