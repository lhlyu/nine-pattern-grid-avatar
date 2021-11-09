package group_avatar

import "testing"

func TestNewGroupAvatar(t *testing.T) {
	g := NewGroupAvatar("", []string{
		"http://xx1.jpg",
		"http://xx2.jpg",
	})
	t.Log(g.Call())
}
