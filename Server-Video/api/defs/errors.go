/**
Usage		报错结构定义
Owner 		wsc
StartDate 	20/7/11
UpdateDate	20/7/11
*/
package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"errpr_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct 请求解析失败", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed 用户验证失败", ErrorCode: "002"}}
	ErrorDB                     = ErrResponse{HttpSC: 500, Error: Err{Error: "DB ops failed 数据库处理错误", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal service error session转换失败", ErrorCode: "004"}}
)
