/**
timer -> setup -> start{trigger ->task ->runner}
*/
package taskrunner

import (
	"fmt"
	"time"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) StartWorker() {

	//for c := range w.ticker.C{
	//	有时间损耗，时间长会有延迟
	//}

	for {
		select {
		case <-w.ticker.C:
			fmt.Println(time.Now())
			go w.runner.StartAll()
		}
	}
}

func Start() {
	//Start video file cleaning
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(5*60, r) //一天清理一次
	go w.StartWorker()      //启动一个定时器任务

}
