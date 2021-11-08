package group_avatar

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
	"path"
)

var (
	client *oss.Client
	bucket string
	domain string
)

func init() {
	accessid := os.Getenv("ALIYUN_ACCESS_ID")
	accesssecret := os.Getenv("ALIYUN_ACCESS_SECRET")
	endpoint := os.Getenv("ALIYUN_OSS_VPC_ENDPOINT")
	bucket = os.Getenv("ALIYUN_OSS_VPC_BUCKET")
	domain = os.Getenv("ALIYUN_OSS_URL")

	var err error
	client, err = oss.New(endpoint, accessid, accesssecret)
	if err != nil {
		panic(err)
	}
}

// 下载资源
func downloadObjects(objectKeys []string) []image.Image {
	if len(objectKeys) == 0{
		return nil
	}
	bk,err := client.Bucket(bucket)
	if err != nil{
		log.Println("downObjects err:",err)
		return nil
	}

	imgs := make([]image.Image,0)
	for _, key := range objectKeys {
		fl,err := bk.GetObject(key)
		if err != nil{
			log.Println("downObjects GetObject", key, err)
			continue
		}

		img,_,err := image.Decode(fl)
		fl.Close()
		if err != nil{
			log.Println("downObjects Decode", key, err)
			continue
		}

		imgs = append(imgs, img)

	}
	return imgs
}

// 上传资源
func uploadObject(objectKey string, stream io.Reader) error {
	bk,err := client.Bucket(bucket)
	if err != nil{
		log.Println("uploadObject err:",err)
		return err
	}
	if err := bk.PutObject(objectKey, stream);err != nil{
		log.Println("uploadObject PutObject",objectKey, err)
		return err
	}
	return nil
}

// 获取图片地址
func getImageUrl(objectKey string) string {
	return path.Join(domain, objectKey)
}