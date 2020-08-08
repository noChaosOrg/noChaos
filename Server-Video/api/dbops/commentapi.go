/**
Usage		评论comment的DB操作
Owner 		wsc
History 	20/7/12 wsc:created
*/

package dbops

import (
	"github.com/noChaos1012/noChaos/Server-Video/api/defs"
	"github.com/noChaos1012/noChaos/Server-Video/api/utils"
)

/**
添加新评论
*/
func AddNewComment(vid, content string, aid int) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id,video_id,author_id,content) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

/**
查询评论列表
*/
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id,users.login_name,comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ?
		AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
		ORDER BY comments.time DESC`)
	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{
			Id:         id,
			VideoId:    vid,
			AuthorName: name,
			Content:    content,
		}
		res = append(res, c)
	}

	defer stmtOut.Close()
	return res, err
}
