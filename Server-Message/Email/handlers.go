package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/noChaos1012/noChaos/Message/mail/conduct"
	"github.com/noChaos1012/noChaos/Message/mail/defs"
	"io/ioutil"
	"log"
	"net/http"
)

/**
Title:发送邮件
URL:[POST]http://127.0.0.1:40000/mail/sendMail
Input:Email结构体
Output:处理结果
Editor:wsc
*/
func sendEmail(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {

	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.Email{}
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	err := conduct.SendMail(ubody)
	if err != nil {
		log.Printf("SendMail ERROR :%v", err)
		sendErrorResponse(w, defs.ErrorConduct)
		return
	}
	sendNormalResponse(w, "发送成功", 200)
	return
}
