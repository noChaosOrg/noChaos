/**
流控，防止带宽、内存过载
*/
package main

import "log"

/**
bucket-token 算法
bucket token 库
request 发送一个token
response 注销一个token
从而控制请求总数
× lock
√ channel shared channel instead of shared memory 共享通道同步线程间的信息，而不是共享内存
*/

type ConnLimiter struct {
	concurrentConn int      //个数
	bucket         chan int //通道，传值使用
}

/**
初始化
*/
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc), //buffer channel 异步通道 nobuffer channel 同步会阻塞

	}
}

/**
判断token数量是否超出
*/

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}

/**
释放连接
*/
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connection coming %d", c)
}
