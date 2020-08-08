/**
Usage		用户user的DB操作
Owner 		wsc
History 	20/7/11 wsc:created
*/
package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/noChaos1012/noChaos/Server-Video/api/defs"
	"log"
)

func AddUserCredential(loginName, pwd string) error {
	// dbConn.Prepare 预编译
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	//在队列结束时执行，会有性能损耗 但在跳出情况多时，需要使用defer
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("GetUserCredential error: %s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func GetUser(loginName string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("SELECT * FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("GetUserCredential error: %s", err)
		return nil, err
	}
	var id int
	var pwd, name string
	err = stmtOut.QueryRow(loginName).Scan(&id, &name, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmtOut.Close()
	return &defs.User{
		Id:        id,
		LoginName: name,
		Pwd:       pwd,
	}, nil
}

func DeleteUser(loginName, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
