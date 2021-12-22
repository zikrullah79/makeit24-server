package worker

import (
	"fmt"
	"sync"
	"zikrullah79/makeit24-server/pkg/mongo"
	"zikrullah79/makeit24-server/pkg/mongo/model"
)

type Worker struct {
	ID          uint
	WorkQueue   chan WorkRequest
	WorkerQueue chan chan WorkRequest
	lock        sync.Mutex
	QuitSignal  chan bool
}

type WorkRequest struct {
	WorkType int
	Result   chan interface{}
	Data     interface{}
	Error    error
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.WorkQueue
			select {
			case work := <-w.WorkQueue:
				w.lock.Lock()
				if work.WorkType == 1 {
					d, err := mongo.GetAllScores()
					if err != nil {
						work.Error = err
						continue
					}
					work.Result <- d
				} else if work.WorkType == 2 {
					// in := work.Data.(model.Score)
					d, err := mongo.InsertScore(work.Data.(*model.Score))
					if err != nil {
						work.Error = err
						continue
					}
					work.Result <- d
				} else if work.WorkType < 10 {
					fmt.Printf("Work Type %v, worked by %v\n", work.WorkType, w.ID)
				} else {
					fmt.Print("Work Type else\n")
				}
				w.lock.Unlock()
			case <-w.QuitSignal:
				fmt.Printf("Worker %v Stopped", w.ID)
				return
			}
		}
	}()

}
