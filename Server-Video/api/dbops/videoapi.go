/**
Usage		视频video_info的DB操作
Owner 		wsc
History 	20/7/11 wsc:created
*/
package dbops

import (
	"database/sql"
	"github.com/noChaos1012/noChaos/Server-Video/api/defs"
	"github.com/noChaos1012/noChaos/Server-Video/api/utils"
	"log"
	"time"
)

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//创建UUID
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	//创建时间
	t := time.Now()
	ctime := t.Format("Jan 02 2006,15:04:05") //M D Y,HH:MM:SS

	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info
		(id,author_id,name,display_ctime) VALUES(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		log.Printf("AddNewVideo error: %s", err)
		return nil, err
	}
	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}
	return res, err
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM video_info WHERE id = ?")
	if err != nil {
		return nil, err
	}
	var aid int
	var ctime, name string
	err = stmtOut.QueryRow(vid).Scan(&aid, &ctime, &name)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("GetVideoInfo error: %s", err)
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		log.Printf("DeleteVideoInfo error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

/**
查询所有视频列表
*/

func ListVideoInfo(u string) ([]*defs.VideoInfo, error) {

	stmtOut, err := dbConn.Prepare(`SELECT
video_info.id ,
	video_info.name ,
	video_info.author_id ,
	video_info.display_ctime
FROM
	video_info
INNER JOIN users ON video_info.author_id = users.id
WHERE
	users.login_name = ?
ORDER BY
	video_info.create_time DESC`)

	var res []*defs.VideoInfo
	rows, err := stmtOut.Query(u)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var author_id int
		var id, name, display_ctime string
		if err := rows.Scan(&id, &name, &author_id, &display_ctime); err != nil {
			return res, err
		}
		v := &defs.VideoInfo{
			Id:           id,
			AuthorId:     author_id,
			Name:         name,
			DisplayCtime: display_ctime,
		}
		res = append(res, v)
	}

	defer stmtOut.Close()
	return res, err
}
