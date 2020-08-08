package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/noChaos1012/noChaos/Server-Video/api/dbops"
	"github.com/noChaos1012/noChaos/Server-Video/api/defs"
	"github.com/noChaos1012/noChaos/Server-Video/api/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func AddNewVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	nvbody := &defs.NewVideo{}
	if err := json.Unmarshal(res, nvbody); err != nil {
		log.Printf("err :%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	vi, err := dbops.AddNewVideo(nvbody.AuthorId, nvbody.Name)
	if err != nil {
		log.Printf("Error in AddNewVide:%s", err)
		sendErrorResponse(w, defs.ErrorDB)
		return
	}

	if resp, err := json.Marshal(&vi); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusCreated)
	}
	return
}

//TODO 分页查询

func ListAllVideos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}
	uname := ps.ByName("username")
	//0, utils.GetCurrentTimestampSec()
	vs, err := dbops.ListVideoInfo(uname)
	if err != nil {
		log.Printf("Error in ListAllVideos:%s", err)
		sendErrorResponse(w, defs.ErrorDB)
		return
	}
	videos := &defs.VideosInfo{Videos: vs}
	if resp, err := json.Marshal(&videos); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusOK)
	}

}

func DeleteVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}
	vid := ps.ByName("vid-id")
	err := dbops.DeleteVideoInfo(vid)
	if err != nil {
		log.Printf("Error in DeleteVideo : %s", err)
		sendErrorResponse(w, defs.ErrorDB)
		return
	}
	utils.SendDeleteVideoRequest(vid)
	sendNormalResponse(w, "", http.StatusNoContent)
}
