package group_avatar

import "image"

// 九种数量头像的展示坐标
var points = [9][]image.Point{
	{
		{user_avatar_size + avatar_gap, user_avatar_size + avatar_gap},
	},
	{
		{group_avatar_size / 2 - avatar_gap / 2 - user_avatar_size, user_avatar_size + avatar_gap},
		{group_avatar_size / 2 + avatar_gap / 2, user_avatar_size + avatar_gap},
	},
	{
		{(group_avatar_size - user_avatar_size) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},
		{(group_avatar_size - 2*user_avatar_size - avatar_gap) / 2, (group_avatar_size + avatar_gap) / 2},
		{(group_avatar_size - avatar_gap ) / 2, (group_avatar_size + avatar_gap) / 2},
	},
	{
		{(group_avatar_size - user_avatar_size * 2 - avatar_gap) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},
		{(group_avatar_size + avatar_gap) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},
		{(group_avatar_size - 2*user_avatar_size - avatar_gap) / 2, (group_avatar_size + avatar_gap) / 2},
		{(group_avatar_size + avatar_gap ) / 2, (group_avatar_size + avatar_gap) / 2},
	},
	{
		{(group_avatar_size - user_avatar_size * 3 - avatar_gap) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},
		{(group_avatar_size - user_avatar_size) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},
		{(group_avatar_size + user_avatar_size + avatar_gap) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},
		{(group_avatar_size - 2*user_avatar_size - avatar_gap) / 2, (group_avatar_size + avatar_gap) / 2},
		{(group_avatar_size - avatar_gap ) / 2, (group_avatar_size + avatar_gap) / 2},
	},
	{
		{(group_avatar_size - user_avatar_size * 3 - avatar_gap) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},
		{(group_avatar_size - user_avatar_size) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},
		{(group_avatar_size + user_avatar_size + avatar_gap) / 2, (group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2},

		{(group_avatar_size - user_avatar_size * 3 - avatar_gap) / 2, (group_avatar_size + avatar_gap) / 2},
		{(group_avatar_size - user_avatar_size) / 2, (group_avatar_size + avatar_gap) / 2},
		{(group_avatar_size + user_avatar_size + avatar_gap) / 2, (group_avatar_size + avatar_gap) / 2},
	},
	{
		{user_avatar_size + avatar_gap, 0},

		{0, user_avatar_size + avatar_gap},
		{user_avatar_size + avatar_gap, user_avatar_size + avatar_gap},
		{2 * user_avatar_size + 2 * avatar_gap, user_avatar_size + avatar_gap},

		{0, 2 * user_avatar_size + 2 * avatar_gap},
		{user_avatar_size + avatar_gap, 2 * user_avatar_size + 2 * avatar_gap},
		{2 * user_avatar_size + 2 * avatar_gap, 2 * user_avatar_size + 2 * avatar_gap},
	},
	{
		{(group_avatar_size - 2 * user_avatar_size - avatar_gap) / 2 , 0},
		{(group_avatar_size + avatar_gap) / 2, 0},

		{0, user_avatar_size + avatar_gap},
		{user_avatar_size + avatar_gap, user_avatar_size + avatar_gap},
		{2 * user_avatar_size + 2 * avatar_gap, user_avatar_size + avatar_gap},

		{0, 2 * user_avatar_size + 2 * avatar_gap},
		{user_avatar_size + avatar_gap, 2 * user_avatar_size + 2 * avatar_gap},
		{2 * user_avatar_size + 2 * avatar_gap, 2 * user_avatar_size + 2 * avatar_gap},
	},
	{
		{0, 0},
		{user_avatar_size + avatar_gap, 0},
		{2 * user_avatar_size + 2 * avatar_gap, 0},

		{0, user_avatar_size + avatar_gap},
		{user_avatar_size + avatar_gap, user_avatar_size + avatar_gap},
		{2 * user_avatar_size + 2 * avatar_gap, user_avatar_size + avatar_gap},

		{0, 2 * user_avatar_size + 2 * avatar_gap},
		{user_avatar_size + avatar_gap, 2 * user_avatar_size + 2 * avatar_gap},
		{2 * user_avatar_size + 2 * avatar_gap, 2 * user_avatar_size + 2 * avatar_gap},
	},
}

// 计算每个用户头像的位置
func (g *groupAvatar) getPoints() bool {
	g.points = points[len(g.images) - 1]
	return true
}