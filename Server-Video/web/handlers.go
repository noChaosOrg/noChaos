package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"nochaos/config"
)

type IndexPage struct {
	Name string
}

/**
* URL		[GET/POST]/
* name		进入主页
* desc
* input		/
* output	/
* history	20/7/16 22/00 wsc: created
 */

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		page := &IndexPage{Name: "王速超"}
		t, e := template.ParseFiles("./templates/index.html")
		if e != nil {
			log.Printf("Parsing template index.html error:%s", e)
			return
		}
		e = t.Execute(w, page)
		if e != nil {
			log.Printf("Execute Page index.html error:%s", e)
			return
		}
		return
	}

	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		//如果不为空则重定向到home
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

}

type HomePage struct {
	Name string
}

/**
* URL		[GET/POST]/home
* name		进入用户主页
* desc
* input		/
* output	/
* history	20/7/16 22/00 wsc: created
 */

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		//如果不为空则重定向到home
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")
	var page *HomePage
	if len(cname.Value) != 0 {
		page = &HomePage{Name: cname.Value}
	} else if len(fname) != 0 {
		page = &HomePage{Name: fname}
	}
	t, e := template.ParseFiles("./templates/home.html")
	if e != nil {
		log.Printf("Parsing template home.html error:%s", e)
		return
	}
	e = t.Execute(w, page)
	if e != nil {
		log.Printf("Execute Page home.html error:%s", e)
		return
	}
	return
}

/**
* URL		[POST]/api
* name		接口透传请求
* desc
* input		ApiBody{}
* output	/
* history	20/7/16 22/00 wsc: created
 */

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Printf("receive request : %v,%v \n", r.Method, r.URL)
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}
	request(apibody, w, r)
	defer r.Body.Close()

}

/**
* URL		[POST]/videos/:vid-id
* name		proxy透传请求
* desc		重置域名
* input		/
* output	/
* history	20/7/16 22/00 wsc: created
 */

func proxyUploadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//res, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("proxy request is %v", r.URL)
	fmt.Printf("proxy request is %v", r.RemoteAddr)
	fmt.Printf("proxy request is %v", r.Header)
	//urlstr := fmt.Sprintf("%v", r.URL)
	u, _ := url.Parse("http://" + config.GetLbAddr() + ":30001")
	log.Printf("\nredirect to %v", u.String())

	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

func proxyVideoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://" + config.GetLbAddr() + ":30001")
	log.Printf("redirect to %v:30001", config.GetLbAddr())
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

/**
* URL		[POST]/dig
* name		接口透传请求
* desc
* input		ApiBody{}
* output	/
* history	20/7/16 22/00 wsc: created
 */

func digHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Printf("receive request : %v,%v \n", r.Method, r.URL)
	//if r.Method != http.MethodPost {
	//	re, _ := json.Marshal(ErrorRequestNotRecognized)
	//	io.WriteString(w, string(re))
	//	return
	//}
	//res, _ := ioutil.ReadAll(r.Body)
	//apibody := &ApiBody{}
	//if err := json.Unmarshal(res, apibody); err != nil {
	//	re, _ := json.Marshal(ErrorRequestBodyParseFailed)
	//	io.WriteString(w, string(re))
	//	return
	//}
	io.WriteString(w, "dig success")

	//request(apibody, w, r)
	//defer r.Body.Close()

}
