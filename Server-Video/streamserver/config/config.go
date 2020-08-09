package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//配置信息
type Configuration struct {
	OssAddr   string `json:"oss_addr"`
	OssBucket string `json:"oss_bucket"`
	OssDir    string `json:"oss_dir"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./conf/oss.json")

	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}
	err := decoder.Decode(configuration)
	fmt.Println(configuration)
	if err != nil {
		panic(err)
	}
}

func GetOssAddr() string {
	return configuration.OssAddr
}

func GetOssBucket() string {
	return configuration.OssBucket
}

func GetOssDir() string {
	return configuration.OssDir
}
