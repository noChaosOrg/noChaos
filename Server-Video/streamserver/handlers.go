package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/noChaos1012/noChaos/Server-Video/streamserver/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/*
		vid := p.ByName("vid-id")
		vl := VIDEO_DIR + "/" + vid
		video, err := os.Open(vl)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
			return
		}
		w.Header().Set("Content-Type", "video/mp4")
		//文件二进制流传输文件
		http.ServeContent(w, r, "", time.Now(), video) //mod time :io.ReadSeeker
		defer video.Close()
	*/

	//cdn.nochaos.top
	//使用OSS访问
	targetUrl := config.GetOssAddr() + p.ByName("vid-id")
	http.Redirect(w, r, targetUrl, http.StatusMovedPermanently)

}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	log.Printf("uploading video :%v", r.Body)

	//TODO 上传视频处理
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE) //限制最大的读取的最大尺寸
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Printf("Error when try to open file:%v", err)
		sendErrorResponse(w, http.StatusBadRequest, "File is too large") //	超过规定文件大小
		return
	}
	file, _, err := r.FormFile("file") //<form name="file"> 客户端注意配置一致
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666) //路径、数据、权限
	if err != nil {
		log.Printf("Write file error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	//上传到七牛云OSS
	ossRet := UploadToOSS(fn, VIDEO_DIR+fn, config.GetOssBucket())
	if !ossRet {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	os.Remove(VIDEO_DIR + fn)

	w.WriteHeader(http.StatusCreated)
	_, err = io.WriteString(w, "Uploading Successfully")
	if err != nil {
		log.Printf("Return msg error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
}
