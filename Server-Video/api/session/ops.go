/**
Usage		Session的相关操作
Owner 		wsc
History 	20/7/12 wsc:created
*/
package session

import (
	"github.com/noChaos1012/noChaos/Server-Video/api/dbops"
	"github.com/noChaos1012/noChaos/Server-Video/api/defs"
	"github.com/noChaos1012/noChaos/Server-Video/api/utils"
	"sync"
	"time"
)

//支持并发读写，线程安全
var sessionMap *sync.Map

//TODO sync.Map 了解
//TODO nosql 改造

func init() {
	sessionMap = &sync.Map{}
}

/**
* name		获取当前时间
* input		sessionId
* output	当前时间戳
* history	20/7/12 23/00 wsc:created
 */

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

//从数据库中读取Session
/**
* name		生成新的Session
* input		用户名
* output	sessionId
* history	20/7/12 23/00 wsc: created
 */
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(key, value interface{}) bool {
		//将查找的数据写入全局变量
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

/**
* name		生成新的Session
* input		用户名
* output	sessionId
* history	20/7/12 23/00 wsc: created
 */
func GenerateNewSessionnId(un string) string {
	id, _ := utils.NewUUID()
	ttl := nowInMilli() + 30*60*1000 //有效时间范围 30min
	ss := &defs.SimpleSession{
		Username: un,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, un, ttl)
	return id
}

/**
* name		Session是否过期
* input		sessionId
* output	sessionId,是否过期
* history	20/7/12 23/00 wsc:created
 */
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	ct := nowInMilli()

	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	} else {
		ss, err := dbops.RetrieveSession(sid)
		if err != nil || ss == nil {
			return "", true
		}

		if ss.TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		sessionMap.Store(sid, ss)
		return ss.Username, false
	}
	return "", true

}
