package taskrunner

import (
	"fmt"
	"github.com/noChaos1012/noChaos/Server-Video/scheduler/config"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var AK, SK string

//VIDEO_BUCKET

func init() {
	//七牛云密钥信息
	AK = "qzgEtUpjZbtbFj8TKNuWi0xtmUI7qlm9tx_ZAZbV"
	SK = "g00Xbrc1XpW56I6RIg24uMrNmRMSHbeNWTbuYoyC"
}

//OSS 文件删除
func deleteOSSFile(filename string) error {

	mac := qbox.NewMac(AK, SK)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	bucketManager := storage.NewBucketManager(mac, &cfg)

	fileInfo, err := bucketManager.Stat(config.GetOssBucket(), filename)
	if err != nil {
		fmt.Printf("\nFind File On OSS Failed，查找OSS文件失败,error:%v", err) //查找不到文件
		return err
	}
	fmt.Println(fileInfo.String())
	err = bucketManager.Delete(config.GetOssBucket(), filename)
	if err != nil {
		fmt.Printf("\nDelelte File On OSS Failed,删除OSS文件失败,error:%v", err)
		return err
	}
	return nil
}
