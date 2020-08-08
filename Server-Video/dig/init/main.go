package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var ualist = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:6.0) Gecko/20100101 Firefox/6.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 2.0.50727; SLCC2; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; Tablet PC 2.0; .NET4.0E)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; InfoPath.3)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; GTB7.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)",
}

// 资源
type resource struct {
	url    string
	target string
	start  int
	end    int
}

func ruleResource() []resource {
	var res []resource
	r1 := resource{
		url:    "http://localhost:8080/",
		target: "",
		start:  0,
		end:    0,
	}

	r2 := resource{
		url:    "http://localhost:8080/home.html",
		target: "",
		start:  0,
		end:    0, //数量
	}
	res = append(res, r1, r2)
	return res
}

func buildURL(res []resource) []string {
	var list []string
	for _, resItem := range res {
		if len(resItem.target) == 0 {
			list = append(list, resItem.url)
		} else {
			for i := resItem.start; i < resItem.end; i++ {
				url := strings.Replace(resItem.url, resItem.target, strconv.Itoa(i), -1)
				list = append(list, url)
			}
		}
	}
	return list
}

func makeLog(current, refer, ua string) string {
	u := url.Values{}
	u.Set("time", time.Now().String())
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramStr := u.Encode()
	logTemplate := `127.0.0.1 - - [07/Aug/2020:00:53:47 +0800] "GET /dig?{$paramStr} HTTP/1.1" 200 43 "http://127.0.0.1:8080/" "{$ua}" "-"`

	logTemplate = strings.Replace(logTemplate, "{$paramStr}", paramStr, -1)
	return strings.Replace(logTemplate, "{$ua}", ua, -1)
}

//获取范围内随机整数
func randInt(start, end int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if start > end {
		return end
	}
	return r.Intn(end-start) + start
}

func main() {
	total := flag.Int("total", 100, "创建行数")
	filePath := flag.String("filePath", "/log/dig.log", "文件路径")
	clear := flag.String("clear", "N", "是否清空")

	flag.Parse()

	res := ruleResource()
	list := buildURL(res)

	var logStack string
	for i := 0; i < *total; i++ {
		current := list[randInt(0, len(list)-1)]
		refer := list[randInt(0, len(list)-1)]
		ua := ualist[randInt(0, len(ualist)-1)]
		logStr := makeLog(current, refer, ua)
		//覆盖写方法
		//ioutil.WriteFile(*filePath, []byte(logStr), 0644)
		logStack += logStr + "\n"
	}
	//添加写文件
	//fmt.Println(logStack + "\n")
	fmt.Printf("total : %d\n", *total)
	fmt.Println("filePath : " + *filePath)
	fmt.Println("clear : " + *clear)
	if *clear == "Y" {
		ioutil.WriteFile(*filePath, []byte(logStack), 0644)
	} else {
		fd, _ := os.OpenFile(*filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		fmt.Printf("%T ,%v", fd, fd)
		cols, err := fd.Write([]byte(logStack))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(cols)
		fd.Close()
	}
	fmt.Println("done.\n")
}
