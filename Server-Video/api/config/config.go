package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//配置信息
type Configuration struct {
	LBAddr string `json:"lb_addr"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./conf/lb.json")

	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}
	err := decoder.Decode(configuration)
	fmt.Println(configuration)
	if err != nil {
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LBAddr
}
