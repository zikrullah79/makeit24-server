package worker

import "fmt"

type Worker struct {
	ID          uint
	WorkQueue   chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitSignal  chan bool
}

type WorkRequest struct {
	WorkType int
	Data     interface{}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.WorkQueue
			select {
			case work := <-w.WorkQueue:
				if work.WorkType < 10 {
					fmt.Printf("Work Type %v, worked by %v\n", work.WorkType, w.ID)
				} else {
					fmt.Print("Work Type else\n")
				}
			case <-w.QuitSignal:
				fmt.Printf("Worker %v Stopped", w.ID)
				return
			}
		}
	}()

}
