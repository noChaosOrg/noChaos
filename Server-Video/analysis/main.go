package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//终端输入参数
type cmdParams struct {
	logFilePath string
	routineNum  int
}

type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

//数据传输
type urlData struct {
	data digData
	uid  string //用户id
}

//存储
type urlNode struct {
}

//
type storageBlock struct {
	counterType  string  //统计类型 pv/uv
	storageModel string  //存储类型
	unode        urlNode //存储数据
}

//日志
var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	//获取参数
	logFilePath := flag.String("logFilePath", "/log/dig.log", "日志路径")
	routineNum := flag.Int("routineNum", 5, "协程数量")
	l := flag.String("l", "/tmp/analysisLog", "日志保存路径")

	flag.Parse()
	params := cmdParams{
		logFilePath: *logFilePath,
		routineNum:  *routineNum,
	}
	//打日志
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = logFd
		defer logFd.Close()
	} else {
		panic(err)
	}

	log.Infoln("Exec start.")
	log.Infoln("Params:logFilePath=%s,routineNum=%d,l=%s", *logFilePath, *routineNum, *l)

	//初始化channel用于数据传递
	var logChannel = make(chan string, 2*params.routineNum) //流量大，双倍数量
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)
	var storageChannel = make(chan storageBlock, params.routineNum)

	//创建日志消费者
	go readFileByLine(params, logChannel)

	//创建一组日志处理
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	//创建PV、UV统计器

	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)

	//创建存储器
	go dataStorage(storageChannel)
	time.Sleep(100 * time.Second)
}

func dataStorage(storageChannel chan storageBlock) {

}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {

}

func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock) {

}

func logConsumer(logChannel chan string, pvChannel, uvChannel chan urlData) {

}

func readFileByLine(params cmdParams, logChannel chan string) {

}
