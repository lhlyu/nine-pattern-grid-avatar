package group_avatar

import (
	"github.com/nfnt/resize"
	"image"
)

// 缩小图片尺寸
func (g *groupAvatar) scale() bool {
	newImages := make([]image.Image,0)
	for _, img := range g.images {
		newImg := resize.Resize(user_avatar_size, user_avatar_size, img, resize.Lanczos3)
		if newImg == nil{
			continue
		}
		newImages = append(newImages, newImg)
	}
	if len(newImages) == 0{
		return false
	}
	g.images = newImages
	return true
}