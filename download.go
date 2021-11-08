package group_avatar

import (
	"net/url"
	"strings"
)

// 下载图片
func (g *groupAvatar) downloadImages() bool {
	// 处理图片，只拿path
	objectKeys := make([]string,0)
	for _, img := range g.imgs {
		u,_ := url.Parse(img)
		if u.Path == ""{
			continue
		}
		objectKeys = append(objectKeys, strings.TrimLeft(u.Path,"/"))
	}

	// 下载
	images := downloadObjects(objectKeys)
	if len(images) == 0{
		return false
	}
	g.images = images
	return true
}