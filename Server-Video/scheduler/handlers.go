package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/noChaos1012/noChaos/Server-Video/scheduler/dbops"
	"net/http"
)

/**
* URL		[GET]/video-delete-record/vid-id
* name		删除视频文件
* desc		使用goroutine 进行很多条记录时的删除和存储
* input		vid-id(文件名)
* output	/
* history	20/7/15 22/00 wsc: created
 */

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")
	return
}
