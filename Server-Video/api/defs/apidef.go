/**
Usage		api的结构定义
Owner 		wsc
StartDate 	20/7/11
UpdateDate	20/7/11
*/
package defs

//输入信息
type UserSession struct {
	Username  string `json:"username"`
	SessionId string `json:"session_id"`
}

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type NewVideo struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
}

type NewComment struct {
	AuthorId int    `json:"author_id"`
	Content  string `json:"content"`
}

//返回结构
type SignUp struct { //&SignIn
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type UserInfo struct {
	Id int `json:"id"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

//Data Model
//用户信息
type User struct {
	Id        int
	LoginName string
	Pwd       string
}

//视频信息
type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

//评论信息
type Comment struct {
	Id         string `json:"id"`
	VideoId    string `json:"video_id"`
	AuthorName string `json:"author_name"`
	Content    string `json:"content"`
}

//简单Session
type SimpleSession struct {
	Username string //login name
	TTL      int64  //识别过期
}
