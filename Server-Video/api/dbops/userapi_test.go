package dbops

import "testing"

// 单元测试步骤 ：init(dblogin,truncate tables) -> run tests ->clear data(truncate tables)

func clearTables() {
	//truncate DDL隐式提交，删节、清空表数据，并不记录操作记录，无法通过回滚恢复
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")

}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	//clearTables()
}

//测试用户操作流 新增->获取->删除->重新获取
//func TestUserWorkFlow(t *testing.T) {
//	t.Run("Add", testAddUser)
//	t.Run("Get", testGetUser)
//	t.Run("Del", testDeleteUser)
//	t.Run("Reget", testRegetUser)
//}

func testAddUser(t *testing.T) {
	err := AddUserCredential("waschild", "123")
	if err != nil {
		t.Errorf("Error of AddUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("waschild")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("waschild", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("waschild")
	if err != nil {
		t.Errorf("Error of ReGetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting User Failed")
	}
}
