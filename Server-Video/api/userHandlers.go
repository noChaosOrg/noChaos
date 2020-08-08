/**
Usage		用户处理实现
Owner 		wsc
StartDate 	20/7/11
UpdateDate	20/7/11

处理流程 handler -> validation{1.request,2.user} -> business logic -> response
*/

package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/noChaos1012/noChaos/Server-Video/api/dbops"
	"github.com/noChaos1012/noChaos/Server-Video/api/defs"
	"github.com/noChaos1012/noChaos/Server-Video/api/session"
	"io/ioutil"
	"log"
	"net/http"
)

/**
* URL		[POST]/login/username
* name		创建用户
* desc		/
* input		/
* output	/
* history	20/7/11 13/00 wsc: created
 */
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取Request的body内容
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		//请求解析错误
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDB)
		return
	}

	id := session.GenerateNewSessionnId(ubody.Username)
	su := &defs.SignUp{
		Success:   true,
		SessionId: id,
	}
	if resp, err := json.Marshal(&su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
		return
	}

}

/**
* URL		[POST]/login/:user_name
* name		登录
* desc		用户登录
* input		/
* output	/
* history	20/7/11 13/00 wsc: created
 */
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("received request body :%s", res)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		log.Printf("err :%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	uname := p.ByName("username")
	if uname != ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	pwd, err := dbops.GetUserCredential(ubody.Username)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser) //密码不符
		return
	}
	id := session.GenerateNewSessionnId(ubody.Username)
	si := &defs.SignUp{
		Success:   true,
		SessionId: id,
	}
	if resp, err := json.Marshal(&si); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)

	} else {
		sendNormalResponse(w, string(resp), http.StatusOK)
	}
}

//获取用户ID
func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unathoried user \n")
		return
	}

	uname := p.ByName("username")
	u, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo:%s", err)
		sendErrorResponse(w, defs.ErrorDB)
		return
	}

	ui := &defs.UserInfo{Id: u.Id}
	if resp, err := json.Marshal(&ui); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusOK)
	}
}
