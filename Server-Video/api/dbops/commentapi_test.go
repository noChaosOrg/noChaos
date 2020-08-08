package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
	clearTables()
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "俺喜欢这个视频"
	err := AddNewComment(vid, content, aid)
	if err != nil {
		t.Errorf("Error of AddComments:%v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514716480 //数据库字段时间戳精度 10位
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments(vid, from, to)
	if err != nil || len(res) == 0 {
		t.Errorf("Error of ListComments:%v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
