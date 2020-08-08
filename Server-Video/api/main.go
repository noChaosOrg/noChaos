package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/noChaos1012/noChaos/Server-Video/api/session"
	"log"
	"net/http"
)

/**
Name:中间件处理结构体
*/
type middleWareHandler struct {
	r *httprouter.Router
}

/**
Name:构建中间件处理
*/
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

/**
Name:实现服务接口，成为http.Handler对象
*/
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	log.Printf("api server receive :%s \n url:%v method:&v\n", r.Header.Get(HEADER_FIELD_UNAME), r.URL.String(), r.Method)
	m.r.ServeHTTP(w, r)
}

/**
Name:注册普通处理路由
*/
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)
	router.GET("/user/:username", GetUserInfo)
	//video&comment相关
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/videos/:vid-id/comments", PostComment)
	router.GET("/videos/:vid-id/comments", ShowComments)

	return router
}

//数据库Session预加载
func Prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	err := http.ListenAndServe(":30000", mh) //listen -> registerhanders -> handlers(goroutine处理)
	if err != nil {
		panic(err)
	}
}

//请求处理顺序
//main -> middleware -> defs(message,err)->handlers ->dbops ->responee
