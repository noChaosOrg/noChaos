/**
//runner -> startDispatcher -> control/data channel
*/

package taskrunner

import "fmt"

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int  //数据长度 待用
	longlived  bool //是否长期存活->是否回收资源
	Dispatcher fn   //调度
	Executor   fn   //执行 闭包
}

func NewRunner(size int, longlived bool, d, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1), //非阻塞
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longlived:  longlived,
		Dispatcher: d, //消费
		Executor:   e, //生产
	}
}

/**
启用调度服务
*/
func (r *Runner) startDispatch() {
	defer func() {
		if !r.longlived { //如果不是长期存活，就回收资源
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data) //写入
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}

			}
			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return //退出
			}
		default:

		}
	}
}

func (r *Runner) StartAll() {
	fmt.Println("准备清理")
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
