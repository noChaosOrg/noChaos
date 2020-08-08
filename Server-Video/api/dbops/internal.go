/**
Usage		Session相关的DB操作
Owner 		wsc
History 	20/7/12 wsc:created
*/

package dbops

import (
	"database/sql"
	"github.com/noChaos1012/noChaos/Server-Video/api/defs"
	"log"
	"strconv"
	"sync"
)

/**
* name		新增Session记录
* input		sessionId,用户名,时间标识
* output	error
* history	20/7/12 18/00 wsc: created
 */
func InsertSession(sid, uname string, ttl int64) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id,TTL,login_name) VALUE (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return err
}

/**
* name		查询Session信息
* input		sessionId
* output	session , error
* history	20/7/12 21/00 wsc: created
 */

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TLL,login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}

	var ttl, uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return nil, nil
}

/**
* name		查询全部Session
* input		/
* output	session键值对 ,err
* history	20/7/12 21/00 wsc: created
 */

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id, ttlstr, login_name string
		if err = rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("restrieve sessions error:%s", err)
			break
		}

		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err != nil {
			ss := &defs.SimpleSession{
				Username: login_name,
				TTL:      ttl,
			}
			m.Store(id, ss)
			log.Printf("session id:%s,ttl:%d", id, ss.TTL)
		}
	}
	return m, err

}

/**
* name		删除Session
* input		sessionId
* output	error
* history	20/7/12 21/00 wsc: created
 */
func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id= ?")
	if err != nil {
		log.Printf("DeleteSession error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(sid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
