/**
Usage		Session校验
Owner 		wsc
History 	20/7/13 wsc:created
*/
package main

import (
	"github.com/noChaos1012/noChaos/Server-Video/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

/**
Name:验证用户session
*/
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	//过期
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

/**
Name:验证用户
Note:IAM 管理用户权限，区分普通用户与后台用户 SSO/Rbac(role based access control)
*/

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		//TODO 发送错误信息
		//sendErrorResponse(w)
		return false
	}
	return true
}
