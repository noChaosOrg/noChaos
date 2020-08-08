/**
api->video_id - mysql ->
dispatcher ->mysql-video_id -> datachannel ->
executor -> datachannel ->video_id ->delete video
*/

package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//批量读取待删除视频记录
func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOuts, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")
	var ids []string
	if err != nil {
		return ids, err
	}

	rows, err := stmtOuts.Query(count)
	if err != nil {
		//TODO 日志分级
		log.Printf("Query VideoDeletionRecord error:%v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err = rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}

	defer stmtOuts.Close()

	return ids, err
}

//删除视频记录
func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id = ?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("DelVideoDeletionRecord error:%v", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}
