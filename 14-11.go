package main

const (
	AvailableMemory         = 10 << 20                                  //10 MB
	AverageMemoryPerRequest = 10 << 10                                  //10 KB
	MAXREQS                 = AvailableMemory / AverageMemoryPerRequest //here amounts to 1000
)

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {}

func handle(r *Request) {
	process(r)
	<-sem
}

func Server(queue chan *Request) {
	for {
		sem <- 1
		request := <-queue
		go handle(request)
	}
}

func main() {
	queue := make(chan *Request)
	go Server(queue)
}
