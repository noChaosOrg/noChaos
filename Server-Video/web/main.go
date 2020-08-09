package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.POST("/", indexHandler)
	router.GET("/home", homeHandler)
	router.POST("/home", homeHandler)
	router.POST("/api", apiHandler) //透传api
	router.GET("/dig", digHandler)  //打点监听
	router.GET("/videos/:vid-id", proxyVideoHandler)
	router.POST("/upload/:vid-id", proxyUploadHandler)

	//文件路径转发
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}

func main() {
	r := RegisterHandler()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}

//cross origin resource sharing 跨域资源共享 -> proxy 透传
