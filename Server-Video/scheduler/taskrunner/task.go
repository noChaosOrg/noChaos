package taskrunner

import (
	"errors"
	"fmt"
	"github.com/noChaos1012/noChaos/Server-Video/scheduler/dbops"
	"log"
	"sync"
)

func deleteVideo(vid string) error {
	//删除本地文件
	//err := os.Remove(VIDEO_PATH + vid)
	//
	////文件存在且报错
	//if err != nil && !os.IsNotExist(err) {
	//	log.Printf("Deleting video error:%v", err)
	//	return err
	//}
	//
	//return err

	//删除OSS文件
	err := deleteOSSFile(vid)
	//文件存在且报错
	if err != nil && string(err.Error()) != "no such file or directory" {
		log.Printf("Deleting video error:%v", err)
		return err
	}
	return err
}

/**
功能：读取需要清理的文件信息
描述：读取数据库中数据，并往data channel里写数据
*/

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error:%v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}
	fmt.Printf("需要删除数据:%v", res)

	for _, id := range res {
		fmt.Println("读取，写入channel:", id)
		dc <- id
	}

	return nil
}

/**
功能：删除清理文件
描述：监听data channel数据写入，则执行删除
*/
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{} //线程安全
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			wg := sync.WaitGroup{} //等待go func完成后再继续执行后续操作 ，否则此处会有重复读取删除现象
			wg.Add(1)
			go func(vid interface{}) {
				if err := deleteVideo(vid.(string)); err != nil {
					errMap.Store(vid, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(vid.(string)); err != nil {
					errMap.Store(vid, err)
					return
				}
				wg.Done()
			}(vid)
			wg.Wait()

		default:
			break forloop
		}
	}

	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
