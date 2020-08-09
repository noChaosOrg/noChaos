/**
Usage		工具包
Owner 		wsc
History 	20/7/11 wsc:created +NewUUID()
*/
package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/noChaos1012/noChaos/Server-Video/api/config"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:12], uuid[12:16]), err
}

//获取当前时间：秒级
func GetCurrentTimestampSec() int {
	time, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	return time
}
func SendDeleteVideoRequest(vid string) {
	addr := config.GetLbAddr() + ":30002"
	url := "http://" + addr + "/video-delete-record/" + vid
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Sending deleting video request error:%v", err)
	}
}
