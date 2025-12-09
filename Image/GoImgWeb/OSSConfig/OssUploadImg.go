package OSSConfig

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

func OssuploadImg(path string, reader io.Reader) bool {
	client, err := oss.New(EndPoint, OSS_ACCESS_KEY_ID, OSS_ACCESS_KEY_SECRET)
	if err != nil {
		panic(err)
		return false
	}
	bucket, err := client.Bucket("yliken-images-test")
	if err != nil {
		panic(err)
		return false
	}
	if err = bucket.PutObject(path, reader); err != nil {
		return false
	}
	return true
}
