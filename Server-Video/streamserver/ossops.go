package main

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"log"
)

////获取token
//func getToken()  string{
//	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
//	putPolicy := storage.PutPolicy{
//		Scope: BUCKET,
//	}
//	mac := qbox.NewMac(AK, SK)
//	token := putPolicy.UploadToken(mac)
//	return token
//}

func UploadToOSS(filename, path, bn string) bool {

	putPolicy := storage.PutPolicy{
		Scope: bn,
	}
	mac := qbox.NewMac(OSS_AK, OSS_SK)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan

	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "test",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, filename, path, &putExtra)
	if err != nil {
		log.Printf("Uploading object error:%v", err)
		return false
	}
	fmt.Println(ret.Key, ret.Hash)
	return true
}
