package group_avatar

import (
	"image"
	"image/color"
	"image/draw"
)

// 绘图
func (g *groupAvatar) drawImages() bool {
	// 创建一个矩形画布
	avatar := image.NewRGBA(image.Rect(0, 0, group_avatar_size, group_avatar_size))

	// 填充颜色
	for m := 0; m < group_avatar_size; m++ {
		for n := 0; n < group_avatar_size; n++ {
			avatar.SetRGBA(m, n, color.RGBA{R: 189, G: 195, B: 199, A: 255})
		}
	}

	// 将每个头像绘到指定坐标
	for i, image := range g.images {
		draw.Draw(avatar, avatar.Bounds(), image, image.Bounds().Min.Sub(g.points[i]), draw.Src)
	}

	g.out = avatar
	return true
}
