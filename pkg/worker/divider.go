package worker

import (
	"math/rand"
)

var WorkerQueue chan chan WorkRequest

func newWorker() *Worker {
	return &Worker{
		ID:          uint(rand.Uint32()),
		WorkQueue:   make(chan WorkRequest),
		WorkerQueue: WorkerQueue,
		QuitSignal:  make(chan bool),
	}
}
func StartDivider(nWorker int) {
	WorkerQueue = make(chan chan WorkRequest)
	for i := 0; i < nWorker; i++ {
		newWorker().Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				go func() {
					worker := <-WorkerQueue

					worker <- work
				}()
			default:
				continue
			}
		}
	}()
}
