package main

import (
	"io"
	"net/http"
)

//返回信息
func sendResponse(w http.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
