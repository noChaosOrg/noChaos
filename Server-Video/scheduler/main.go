/**
1、用户通过api传达删除指令
2、通过api service 添加删除记录（video_del_rec）
3、定时器启动，读取删除记录，执行删除video文件
*/

package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/noChaos1012/noChaos/Server-Video/scheduler/taskrunner"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}

func main() {
	/**
	通过channel的定义取值也可以进行阻塞block
	c := make(chan  int)
	go taskrunner.Start()
	<-c
	*/

	go taskrunner.Start()
	r := RegisterHandlers()
	err := http.ListenAndServe(":30002", r)
	if err != nil {
		panic(err)
	}
}
