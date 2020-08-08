package main

import (
	"io"
	"net/http"
)

/**
Name:返回报错信息
*/
func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
