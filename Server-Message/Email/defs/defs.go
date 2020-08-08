package defs

type Email struct {
	ServerHost string   `json:"server_host"` //服务地址
	ServerPort int      `json:"server_port"` //服务端口号
	FromEmail  string   `json:"from_email"`  //发送人邮箱
	FromPwd    string   `json:"from_pwd"`    //发送人密码
	Toers      []string `json:"toers"`       //接收人
	Subject    string   `json:"subject"`     //主题
	Body       string   `json:"body"`        //内容
}
