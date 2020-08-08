package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func registerHandlers() *httprouter.Router {
	r := httprouter.New()
	r.POST("/Email/sendMail", sendEmail)
	return r
}

func main() {
	r := registerHandlers()
	err := http.ListenAndServe(":40000", r)
	if err != nil {
		panic(err)
	}
}
