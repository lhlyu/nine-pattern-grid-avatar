package group_avatar

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"time"
)

const (
	// 每个用户头像的边长
	user_avatar_size = 40
	// 每个头像之间的间隙
	avatar_gap = 0
	// 群头像的边长
	group_avatar_size = user_avatar_size * 3 + avatar_gap * 2

	// 群头像命名规则
	group_avatar_name = "group/avatar/%s_%s.jpg"

)

type groupAvatar struct {
	name string
	// 图片地址
	imgs []string
	images []image.Image
	points []image.Point
	out image.Image
}

func NewGroupAvatar(name string,imgs []string) *groupAvatar {
	return &groupAvatar{
		name: name,
		imgs: imgs,
	}
}

// 生成群头像, 返回地址
func (g *groupAvatar) Call() string {
	if len(g.imgs) == 0{
		return ""
	}
	// 下载图片
	if !g.downloadImages() {
		return ""
	}
	// 缩小尺寸
	if !g.scale() {
		return ""
	}
	// 获取每张图的位置
	if !g.getPoints() {
		return ""
	}
	// 画图
	if !g.drawImages() {
		return ""
	}

	// 上传图片
	objectKey := fmt.Sprintf(group_avatar_name, g.name, time.Now().Format("060102150405"))

	buf := &bytes.Buffer{}

	if err := jpeg.Encode(buf,g.out, nil);err != nil{
		log.Println("Call Encode ",g.name, err)
		return ""
	}

	if err := uploadObject(objectKey, buf);err != nil{
		log.Println("Call uploadObject ",g.name, err)
		return ""
	}

	return getImageUrl(objectKey)
}




