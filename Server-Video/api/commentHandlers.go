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

func PostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	cbody := &defs.NewComment{}
	if err := json.Unmarshal(reqBody, cbody); err != nil {
		log.Printf("err :%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	vid := ps.ByName("vid-id")
	err := dbops.AddNewComment(vid, cbody.Content, cbody.AuthorId)
	if err != nil {
		log.Printf("Error in PostComment:%s", err)
		sendErrorResponse(w, defs.ErrorDB)
		return
	}
	sendNormalResponse(w, "ok", http.StatusCreated)
	return
}

//TODO 分页查询

func ShowComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}

	vid := ps.ByName("vid-id")
	cm, err := dbops.ListComments(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ShowComments:%s", err)
		sendErrorResponse(w, defs.ErrorDB)
		return
	}
	cms := &defs.Comments{Comments: cm}
	if resp, err := json.Marshal(&cms); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusOK)
	}
}
