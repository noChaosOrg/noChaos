package dbops

import "testing"

var tempVid string

//测试用户操作流 新增->获取->删除->重新获取
//func TestVideoWorkFlow(t *testing.T) {
//	clearTables()
//	t.Run("PrepareUser", testAddUser)
//	t.Run("Add", testAddNewVideo)
//	t.Run("Get", testGetVideoInfo)
//	t.Run("Del", testDeleteVideoInfo)
//	t.Run("Reget", testRegetVideoInfo)
//}

func testAddNewVideo(t *testing.T) {
	video, err := AddNewVideo(1, "video-1")
	if err != nil {
		t.Errorf("Error of AddNewVideo:%v", err)
	}
	tempVid = video.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo")
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo:%v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of RegetVideoInfo")
	}
	if video != nil {
		t.Errorf("Error of RegetVideoInfo")

	}
}
